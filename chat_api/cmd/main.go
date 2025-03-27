package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	desc "valera/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatAPIServer
}

func (s *server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	fmt.Println(request.Usernames)
	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (s *server) Delete(ctx context.Context, request *desc.DeleteRequest) (*desc.DeleteResponse, error) {
	fmt.Println(request.Id)
	return &desc.DeleteResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func (s *server) SendMessage(ctx context.Context, request *desc.SendMessageRequest) (*desc.SendMessageResponse, error) {
	fmt.Println(request.Text, request.Timestamp, request.From)
	return &desc.SendMessageResponse{
		Empty: &emptypb.Empty{},
	}, nil
}

func main() {
	list, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatAPIServer(s, &server{})

	log.Printf("server listening at %v", list.Addr())

	if err = s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
