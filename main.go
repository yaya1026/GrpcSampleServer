package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	log.Println("connect!")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetGameReviews(ctx context.Context, in *emptypb.Empty) (*game.GameReviews, error) {

	return &game.GameReviews{
		Games: []*game.GameReview{
			{
				Id:        1,
				Title:     "龍が如く0 誓いの場所",
				Provider:  "SEGA",
				ImageUrl:  "https://images-na.ssl-images-amazon.com/images/I/61ZbS8G1NhL._AC_.jpg",
				Comment:   "素晴らしいストーリー、実写では無いが素晴らしい画像、更に格闘ゲームの要素を始め様々な遊び要素が満載で、色々な楽しみ方があり、まだまだ全てを楽し尽くしたわけではないが、娯楽作品としての映画産業が今後衰退してゆくのではないか？と感じさせる程のゲームだった。",
				User:      &game.User{Id: 1, Name: "yaya"},
				CreatedAt: &timestamppb.Timestamp{},
			},
			{
				Id:        2,
				Title:     "龍が如く極",
				Provider:  "SEGA",
				ImageUrl:  "https://images-na.ssl-images-amazon.com/images/I/71jHL%2Bufy2L._AC_SL1000_.jpg",
				Comment:   "龍が如くシリーズは、人気があるのは知っていましたが、極道モノに興味がなく、今まで避けてきました。が、低価格版のリメイクである今作に興味を持ち、購入。爽快なバトルと、描かれる人間ドラマが良く、ハマってしまいました。続けて極2、0を買うほどに。このシリーズをプレイしたことのない方は、この極1から始めてみてはいかがでしょうか？安いしオススメですよ。",
				User:      &game.User{Id: 1, Name: "yaya"},
				CreatedAt: &timestamppb.Timestamp{},
			},
			{
				Id:        2,
				Title:     "龍が如く極2",
				Provider:  "SEGA",
				ImageUrl:  "https://images-na.ssl-images-amazon.com/images/I/61oZv60y9eL._AC_SL1000_.jpg",
				Comment:   "蒼天堀と神室町のマップが一新されていて、新鮮です。町歩きにより本物感があり、映像も精密になっています。",
				User:      &game.User{Id: 1, Name: "yaya"},
				CreatedAt: &timestamppb.Timestamp{},
			},
			{
				Id:        3,
				Title:     "あつまれ どうぶつの森",
				Provider:  "任天堂",
				ImageUrl:  "https://images-na.ssl-images-amazon.com/images/I/71lpG-B9oDL._AC_SL1133_.jpg",
				Comment:   "時間が溶けてしまうほど面白い",
				User:      &game.User{Id: 1, Name: "yaya"},
				CreatedAt: &timestamppb.Timestamp{},
			},
		},
	}, nil
}

func (s *server) GetGameReview(ctx context.Context, in *game.GameReviewRequest) (*game.GameReview, error) {

	return &game.GameReview{
		Id:        1,
		Title:     "龍が如く0 誓いの場所",
		Provider:  "SEGA",
		ImageUrl:  "https://images-na.ssl-images-amazon.com/images/I/61ZbS8G1NhL._AC_.jpg",
		Comment:   "素晴らしいストーリー、実写では無いが素晴らしい画像、更に格闘ゲームの要素を始め様々な遊び要素が満載で、色々な楽しみ方があり、まだまだ全てを楽しみ尽くしたわけではないが、娯楽作品としての映画産業が今後衰退してゆくのではないか？と感じさせる程のゲームだった。",
		User:      &game.User{Id: 1, Name: "yaya"},
		CreatedAt: &timestamppb.Timestamp{},
	}, nil
}
