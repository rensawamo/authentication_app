package app

import (
	"context"

	pb "github.com/RecepieApp/server/gen/buf/proto"
)

// SayHelloはExampleServiceのSayHello RPCを実装します。
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}


/// ユーザがお知らせを取得するためのRPC
func (s *AnnouncementServer) GetAnnouncement(ctx context.Context, in *pb.GetAnnouncementRequest) (*pb.Announcement, error) {
	// ダミーデータのお知らせを作成
	announcement := &pb.Announcement{
			Id:      1,
			Title:   "サンプルお知らせ",
			Content: "これはサンプルのお知らせです。",
	}
	// お知らせをレスポンスに含める
	return announcement, nil
}