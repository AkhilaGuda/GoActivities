package main

import (
	"context"
	"io"
	"log"

	pb "grpc/routeguide"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// CONNECT TO SERVER
	// using insecure credentials here, meaning no TLS encryption
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	// Create a new client stub — this is used to call the remote methods.
	client := pb.NewRouteGuideClient(conn)

	// 1. Unary RPC
	point := &pb.Point{Latitude: 409146138, Longitude: -746188906} // point message
	feature, err := client.GetFeature(context.Background(), point) // send request to server
	if err != nil {
		log.Fatalf("GetFeature failed: %v", err)
	}
	// server response
	log.Printf("GetFeature: %v", feature)

	// 2. Server streaming RPC - one request multiple responses
	rect := &pb.Rectangle{
		Lo: &pb.Point{Latitude: 0, Longitude: 0},   // lower left corner
		Hi: &pb.Point{Latitude: 10, Longitude: 10}, // upper right corner
	}
	// call ListFeatures, which returns a stream of responses
	stream1, _ := client.ListFeatures(context.Background(), rect)
	for {
		// receive one feature at a time from the server stream
		f, err := stream1.Recv()
		// stream is completed
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListFeatures error: %v", err)
		}
		log.Printf("Feature: %v", f)
	}

	// 3. Client streaming RPC - multiple requests, one response

	// RecordRoute lets client send multiple points to server
	// server responds with summary info
	stream2, _ := client.RecordRoute(context.Background())
	for i := 0; i < 5; i++ {
		// send multiple points to server
		stream2.Send(&pb.Point{Latitude: int32(i), Longitude: int32(i)})
	}
	// close the stream and wait for server's single response
	summary, _ := stream2.CloseAndRecv()
	log.Printf("RouteSummary: %v", summary)

	// 4. Bidirectional streaming RPC - Both sides sends multiple messages

	stream3, _ := client.RouteChat(context.Background())
	// channel used to wait until receiving is complete
	waitc := make(chan struct{})
	// go routine to receive messages from the server
	go func() {
		for {
			// receive RouteNote messages from the server
			note, err := stream3.Recv()
			if err == io.EOF {
				// server closed the stream
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("RouteChat error: %v", err)
			}
			log.Printf("Received note: %v", note)
		}
	}()
	// send few RouteNotes from the client side
	for i := 0; i < 3; i++ {
		stream3.Send(&pb.RouteNote{
			Location: &pb.Point{Latitude: int32(i), Longitude: int32(i)},
			Message:  "Message " + string('A'+i), // Message A, Message B
		})
	}
	// close the sending direction stream
	stream3.CloseSend()
	// wait until the receiving goroutine finishes
	<-waitc
}
