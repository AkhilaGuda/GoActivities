package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "grpc/routeguide"

	"google.golang.org/grpc"
)

// server struct implements all RPC methods of RouteGuide servcie
type server struct {
	pb.UnimplementedRouteGuideServer
	// here we can add additional fields eg: in-memory DB, mutex
}

// 1. Unary RPC : GetFeature
func (s *server) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Printf("GetFeature request: %v,%v", point.Latitude, point.Longitude)
	// create and return a Feature (response) for the given point
	return &pb.Feature{
		Name:     "Example Feature",
		Location: point,
	}, nil
}

// Server-side streaming RPC
// client sends one rectangle -> server sends multiple feature messages
func (s *server) ListFeatures(rect *pb.Rectangle, stream grpc.ServerStreamingServer[pb.Feature]) error {

	// sending 5 feature messages to the client
	for i := 0; i < 5; i++ {
		// create a simple feature for each iteration
		feature := &pb.Feature{
			Name:     "Feature " + string('A'+i),
			Location: &pb.Point{Latitude: rect.Lo.Latitude + int32(i), Longitude: rect.Lo.Longitude + int32(i)},
		}
		// send the feature to the client through the stream
		if err := stream.Send(feature); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 2000)
	}
	return nil
}

// Client-side streaming RPC
// client sends multiple points -> server responds with one route summary
func (s *server) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	// count keeps tracks of how many points client sent
	var count int32
	for {
		// receive a point from the client stream
		point, err := stream.Recv()
		if err == io.EOF {
			// send final summary response and close the stream
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   count,     // total number of received points
				FeatureCount: count / 2, // dummy count for example
				Distance:     100,       // example static distance
				ElapsedTime:  1,         // example static duration
			})
		}
		// some other error while receiving
		if err != nil {
			return err
		}
		// log received coordinates
		log.Printf("RecordRoute received: %v,%v", point.Latitude, point.Longitude)
		// Increment received point counter
		count++
	}
}

// Bidirectional streaming RPC - RouteChat
// client and server exchanges RouteNote messages in real time
func (s *server) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		// receive a message RouteNote from client
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		// log received RouteNote from client
		log.Printf("RouteChat received: %s at %v,%v", note.Message, note.Location.Latitude, note.Location.Longitude)
		// send back a reply to the client
		reply := &pb.RouteNote{
			Message:  "Ack: " + note.Message,
			Location: note.Location,
		}
		// send the acknowledgement message to the client
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	// Listen to TCP port 50051 for incoming gRPC connections
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a new grpc server instance
	grpcServer := grpc.NewServer()
	// Register our RouteGuide service implementation with the server
	pb.RegisterRouteGuideServer(grpcServer, &server{})
	log.Println("Server running on :50051")
	// start serving incoming connections
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
