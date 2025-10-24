package main

import (
	"bufio"
	pb "chat/pkg/proto"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server running on localhost:50051
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	defer connection.Close()
	// Create a buffered reader to read user Input from terminal
	reader := bufio.NewReader(os.Stdin)
	// Create a new gRPC client instance
	client := pb.NewCHATClient(connection)
	// open a bidirectional stream with a server
	stream, _ := client.BidiChatRoom(context.Background())
	// Take user input for username and Id
	fmt.Print("Enter recipient username: ")
	var fromId string
	fmt.Scanln(&fromId)
	// Start a goroutine to continuously receive and display messages from the server
	go func() {
		for {
			// receive messages sent from the server
			streamMessage, err := stream.Recv()
			if err != nil {
				log.Println("Stream closed")
				// exit when stream is closed
				return
			}
			// display received message in sender: message format
			fmt.Printf("%s: %s\n", streamMessage.FromId, streamMessage.Message)
		}

	}()
	// continuously read user input and send messages to the chat room
	for {
		text, _ := reader.ReadString('\n') // read message from console
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		// send message to server to broadcast
		stream.Send(&pb.ChatMessage{FromId: fromId, Message: text})
	}

}
