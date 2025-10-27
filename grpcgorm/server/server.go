package main

import (
	"context"
	"fmt"
	db "grpcgorm/db"
	pb "grpcgorm/grpcgorm"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGRPCGORMServer
	blogRepo db.BlogRepository
}

func (s *server) Create(ctx context.Context, req *pb.CreatePostRequest) (*pb.PostResponse, error) {
	post := db.GenericTable{
		Title:   req.Title,
		Content: req.Content,
	}
	if err := s.blogRepo.CreatePost(&post); err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return &pb.PostResponse{
		Id:      int32(post.ID),
		Title:   post.Title,
		Content: post.Content,
	}, nil
}

func (s *server) GetAll(ctx context.Context, req *pb.Empty) (*pb.GetAllResponse, error) {
	posts, err := s.blogRepo.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}

	var pbPosts []*pb.PostResponse
	for _, post := range posts {
		pbPosts = append(pbPosts, &pb.PostResponse{
			Id:      int32(post.ID),
			Title:   post.Title,
			Content: post.Content,
		})
	}
	return &pb.GetAllResponse{
		Posts: pbPosts,
	}, nil
}

func main() {
	fmt.Println("Data base connected")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGRPCGORMServer(grpcServer, &server{})
	log.Println("grpc server running on port : 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
