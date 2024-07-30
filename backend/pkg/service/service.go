package service

import (
	"context"

	pb "github.com/authentication_app/backend/gen/buf/proto"
)

type HelloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	reply := &pb.HelloReply{
		Message: "認証成功 こんにちは",
	}
	return reply, nil
}
