package app

import (
	pb "github.com/RecepieApp/server/gen/buf/proto"
	"github.com/RecepieApp/server/middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// serverはpb.ExampleServiceServerに対する実装です。
type server struct {
	pb.UnimplementedHelloServiceServer
}
type AnnouncementServer struct {	
	pb.UnimplementedAnnouncementServiceServer
}

func (a *Application) setupGrpcServer() {
	// 認証インターセプターを使用してgRPCサーバーを設定
	interceptor := middleware.AuthInterceptor(a.FireAuth)
	a.GrpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	pb.RegisterHelloServiceServer(a.GrpcServer, &server{})
	pb.RegisterAnnouncementServiceServer(a.GrpcServer, &AnnouncementServer{})
	reflection.Register(a.GrpcServer)
}
