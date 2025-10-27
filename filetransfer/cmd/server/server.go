package main

import (
	pb "filetransfer/pkg/proto"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const uploadDir = "./uploads"

type server struct {
	pb.UnimplementedFileServiceServer
}

func NewServer() *server {
	return &server{}
}

// Upload handles client streaming RPC where client sends file data in chunks
func (s *server) Upload(stream pb.FileService_UploadServer) error {
	log.Println("Upload request received")
	file := NewFile()       // custom helper for creating new file
	var fileSize uint32 = 0 // to track total file size
	// ensures file is closed when an error occurs
	defer func() {
		if file.OutputFile != nil {
			if err := file.OutputFile.Close(); err != nil {
				log.Printf("warning failed to close file : %v\n", err)
			}
		}
	}()
	// Receive stream of data from the client
	for {
		req, err := stream.Recv()
		// If filePath not set yet, initialize it using file name
		if file.FilePath == "" {
			file.SetFile(req.GetFileName(), "./uploads")

		}
		// EOf - client finished sending all chunks
		if err == io.EOF {
			log.Println("File upload complete, sending response to client")
			break
		}
		//  Handle other stream errors
		if err != nil {
			log.Printf("Error receiving stream: %v\n", err)
			return status.Error(codes.Internal, err.Error())
		}
		// Create uploads directory if not exists
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return status.Error(codes.Internal, "failed to create uploads directory: "+err.Error())
		}
		// Initialize the file creation only once
		if file.FilePath == "" {
			err = file.SetFile(req.GetFileName(), uploadDir)
			if err != nil {
				log.Printf("Error creating file : %v\n", err)
				return status.Error(codes.Internal, err.Error())
			}
			log.Printf("created file : %s\n", file.FilePath)
		}
		// Get the data chunk from request
		chunk := req.GetChunk()
		// Increment total file size
		fileSize += uint32(len(chunk))
		log.Printf("received a chunk with size: %d\n", fileSize)
		// Write chunk to output file
		if err := file.Write(chunk); err != nil {
			log.Printf("Error writing chunk: %v\n", err)
			return status.Error(codes.Internal, err.Error())
		}
	}
	/// Extract base fole name from path - removes directory prefix
	fileName := filepath.Base(file.FilePath)
	// Send a final response back to the client after upload completes
	log.Printf("Saved file : %s, size : %d\n", fileName, fileSize)
	return stream.SendAndClose(&pb.FileUploadResponse{FileName: fileName, Size: fileSize})

}

// Download handles server streaming RPC where server sends file data in chunks
func (s *server) Download(req *pb.FileDownloadRequest, stream pb.FileService_DownloadServer) error {
	// Build full path of file to send
	filePath := filepath.Join(uploadDir, req.GetFileName())
	// open file for reading
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %v\n", err)
		return status.Error(codes.Internal, err.Error())
	}
	defer file.Close()
	// create buffer to read chunks
	buffer := make([]byte, 64*1024)
	for {
		// read file data into buffer
		n, err := file.Read(buffer)
		if err == io.EOF {
			log.Println("Completed file reading ")
			break
		}
		if err != nil {
			log.Printf("Error reading  file: %v\n", err)
			return status.Error(codes.Internal, err.Error())
		}
		// send chunk back to client
		if err := stream.Send(&pb.FileDownloadResponse{FileName: req.GetFileName(), Chunk: buffer[:n]}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

	}
	log.Printf("File %s is sent successfully\n", req.GetFileName())
	return nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error :%s ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, NewServer())
	log.Println("Server started on port: 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
