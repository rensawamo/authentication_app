package server

import (
	pb "github.com/authentication_app/backend/gen/buf/proto"
	"github.com/authentication_app/backend/middleware"
	"github.com/authentication_app/backend/pkg/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func SetupGrpcServer(a *app.Application) {
	interceptor := middleware.AuthInterceptor(a.FireAuth)
	a.GrpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	pb.RegisterHelloServiceServer(a.GrpcServer, &Server{})
	reflection.Register(a.GrpcServer)
}
