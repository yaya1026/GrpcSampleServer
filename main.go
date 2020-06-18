package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"grpc.sample.server/generated"
)

type server struct {
	generated.BookServiceClient
}

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	generated.RegisterBookServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Println("connect!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetBooks(ctx context.Context, in *emptypb.Empty) (*generated.Books, error){
	return &generated.Books{Books: []*generated.Book{{Id: 1, Title: "Go", Author: "Yaya", Isbn: "978-4-86501-422-8", Overview: "いいよ"}}}, nil
}