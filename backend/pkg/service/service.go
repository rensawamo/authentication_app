package service

import (
	"context"

	pb "github.com/authentication_app/backend/gen/buf/proto"
)

type AnnouncementServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *AnnouncementServer) GetAnnouncement(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// ダミーデータのお知らせを作成
	reply := &pb.HelloReply{
		Message: "認証成功 こんにちは",
	}
	return reply, nil
}
