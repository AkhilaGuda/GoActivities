package main

import (
	pb "chat/pkg/proto"
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

type StreamManager interface {
	Register(streamId string, stream pb.CHAT_BidiChatRoomServer)
	Deregister(streamId string)
	BroadCast(message *pb.ChatMessage)
	GetStreamCount() int
}
type MapManager struct {
	bidiStreams map[string]pb.CHAT_BidiChatRoomServer
	mu          sync.Mutex
}

type server struct {
	pb.UnimplementedCHATServer
	joinRoomStreams  map[string]pb.CHAT_JoinChatRoomServer
	bidiSreamManager StreamManager
}

func NewMapManager() *MapManager {
	return &MapManager{
		bidiStreams: make(map[string]pb.CHAT_BidiChatRoomServer),
	}
}
func newServer() *server {
	return &server{
		joinRoomStreams:  make(map[string]pb.CHAT_JoinChatRoomServer),
		bidiSreamManager: NewMapManager(),
	}
}
func (m *MapManager) Register(streamId string, stream pb.CHAT_BidiChatRoomServer) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.bidiStreams[streamId] = stream
}
func (m *MapManager) Deregister(streamId string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.bidiStreams, streamId)
}
func (m *MapManager) BroadCast(message *pb.ChatMessage) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, stream := range m.bidiStreams {
		if err := stream.Send(message); err != nil {
			log.Printf("Error sending message to stream: %v", err)
		}
	}
}
func (m *MapManager) GetStreamCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.bidiStreams)
}

func (s *server) PrivateSend(ctx context.Context, in *pb.MessageRequest) (*pb.MessageResponse, error) {
	receiver, ok := s.joinRoomStreams[in.ToId]
	if !ok {
		return &pb.MessageResponse{State: "User Is Unavailable"}, nil
	}
	msg := &pb.ChatMessage{FromId: in.FromId, Message: in.Message}
	receiver.Send(msg)
	// s.clients["123"].Send(&pb.ChatMessage{FromId,Message})
	return &pb.MessageResponse{State: "Delivered"}, nil
}

func (s *server) BidiChatRoom(stream pb.CHAT_BidiChatRoomServer) error {
	// receive first message to get the user ID
	message, err := stream.Recv()
	if err != nil {
		return nil
	}

	clientId := message.FromId
	// Register the stream
	s.bidiSreamManager.Register(clientId, stream)

	// defer registration and user left chat room message broadcast
	defer func() {
		s.bidiSreamManager.Deregister(clientId)
		log.Printf("%s left the chat\n", clientId)

		leftMessage := &pb.ChatMessage{
			FromId:  clientId,
			Message: fmt.Sprintf("%s left the chat", clientId),
		}
		s.bidiSreamManager.BroadCast(leftMessage)
	}()
	// notify other users new user joined
	joinMessage := &pb.ChatMessage{
		FromId:  clientId,
		Message: fmt.Sprintf("%s joined the chat", clientId),
	}
	s.bidiSreamManager.BroadCast(joinMessage)

	log.Printf("%s joined chat \n", clientId)
	// listen for message received from the client and broadcast
	for {
		msg, err := stream.Recv()
		if err != nil {
			return err
		}

		broadcastMsg := &pb.ChatMessage{
			FromId:  clientId,
			Message: msg.Message,
		}
		s.bidiSreamManager.BroadCast(broadcastMsg)
		log.Printf("Message from %s : %s", clientId, msg.Message)

	}

}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error :%s ", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterCHATServer(grpcServer, newServer())
	log.Println("Server started on port: 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %s", err)
	}
}
