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
	connection, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect: ", err)
	}
	defer connection.Close()
	reader := bufio.NewReader(os.Stdin)
	client := pb.NewCHATClient(connection)
	stream, _ := client.BidiChatRoom(context.Background())
	fmt.Print("Enter recipient username: ")
	var fromId string
	fmt.Scanln(&fromId)
	go func() {
		for {
			streamMessage, err := stream.Recv()
			if err != nil {
				log.Println("Stream closed")
				return
			}
			fmt.Printf("%s: %s\n", streamMessage.FromId, streamMessage.Message)
		}

	}()
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		stream.Send(&pb.ChatMessage{FromId: fromId, Message: text})
	}

}
