package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	game "grpc.sample.server/generated"
)

type server struct {
	game.GameServiceClient
}

const port = ":8080"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	game.RegisterGameServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Println("connect!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetGameReviews(ctx context.Context, in *emptypb.Empty) (*game.GameReviews, error) {

	return nil, nil
}

func (s *server) GetGameReview(ctx context.Context, in *game.GameReviewRequest) (*game.GameReview, error) {

	return nil, nil
}
