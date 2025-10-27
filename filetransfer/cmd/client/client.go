package main

import (
	"bufio"
	"context"
	pb "filetransfer/pkg/proto"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// folder where downloaded files will be saved
const downloadDir = "./downloads"

type ClientService struct {
	addr      string               // server address ( localhost:50051)
	filePath  string               // file path to u[;pad]
	batchSize int                  // number of bytes to send per chunk
	client    pb.FileServiceClient // generated gRPC client interfaces
}

// upload function uploads file to the gRPC server using client streaming
func (s *ClientService) upload(ctx context.Context, cancel context.CancelFunc) error {
	// create new upload stream to send file chunks
	stream, err := s.client.Upload(ctx)
	if err != nil {
		return err
	}
	// open file from the given file path
	file, err := os.Open(s.filePath)
	if err != nil {
		return err
	}
	// create a buffer of batchsize bytes to read chunks
	buff := make([]byte, s.batchSize)
	batchNumber := 1
	for {
		// read next chunk from file
		num, err := file.Read(buff)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		// Extract the portion actually read
		chunk := buff[:num]
		// send this chunk to the server with filename
		if err := stream.Send(&pb.FileUploadRequest{FileName: GetFileName(s.filePath), Chunk: chunk}); err != nil {
			return err
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}
	// close the stream and receive the servers final response
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Sent -%v bytes - %s\n", res.GetSize(), res.GetFileName())
	return nil
}

// download function handles to downloads a file from gRPC server using server streaming
func (s *ClientService) download(ctx context.Context, fileName string) error {
	// Request a download stream from the server
	stream, err := s.client.Download(ctx, &pb.FileDownloadRequest{FileName: fileName})
	if err != nil {
		return err
	}
	// create a file locally to save the downloaded content
	downloadPath := filepath.Join(downloadDir, fileName)

	// Create the downloads directory if it doesn’t exist
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create downloads directory: %v", err)
	}

	// Create the output file inside downloads folder
	outFile, err := os.Create(downloadPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %v", err)
	}
	defer outFile.Close()
	var totalBytes int
	// Receive chunks until EOF server finishes sending
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break // download complete
		}
		if err != nil {
			return err
		}

		chunk := res.GetChunk()
		totalBytes += len(chunk)
		// write chunk to output file
		if _, err := outFile.Write(chunk); err != nil {
			return err
		}
	}

	log.Printf("File %s downloaded successfully (%d bytes)\n", fileName, totalBytes)
	return nil
}

// helper to extract only file name from full path (eg. "/home/user/sample.txt" -> "sample.txt")
func GetFileName(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	return parts[len(parts)-1]
}
func main() {
	// Connect to the gRPC server running on localhost:50051
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	defer connection.Close()
	// Create a new gRPC client instance
	client := pb.NewFileServiceClient(connection)
	// Create a buffered reader to read user Input from terminal
	reader := bufio.NewReader(os.Stdin)
	log.Print("enter full file path to upload: ")
	filePath, _ := reader.ReadString('\n')
	filePath = strings.TrimSpace(filePath)
	service := &ClientService{
		addr:      "localhost:50051",
		filePath:  filePath,
		batchSize: 64 * 1024, //64kb
		client:    client,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err = service.upload(ctx, cancel)
	if err != nil {
		log.Fatalf("Upload failed: %v", err)
	}
	log.Print("Enter file name to download: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	err = service.download(ctx, fileName)
	if err != nil {
		log.Fatalf("Download failed: %v", err)
	}

}
