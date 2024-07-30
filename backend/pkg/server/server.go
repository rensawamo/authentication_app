package server

import (
	"context"
	"log"
	"net"

	"github.com/authentication_app/backend/config"
	md "github.com/authentication_app/backend/middleware"

	pb "github.com/authentication_app/backend/gen/buf/proto"
	"github.com/authentication_app/backend/middleware"
	"github.com/authentication_app/backend/model"
	"github.com/authentication_app/backend/pkg/api"
	"github.com/authentication_app/backend/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	pb.UnimplementedHelloServiceServer
}

func SetUpGrpcServer(a *model.Application) {
	interceptor := middleware.AuthInterceptor(a.FireAuth)
	a.GrpcServer = grpc.NewServer(grpc.UnaryInterceptor(interceptor))
	helloServer := &service.HelloServer{}
	pb.RegisterHelloServiceServer(a.GrpcServer, helloServer)
	reflection.Register(a.GrpcServer)
}

func StartGrpcServer(a *model.Application) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server is listening on port 50051")
	if err := a.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func SetUpRestServer(a *model.Application) error {
	ctx := context.Background()

	fireClient, err := config.GetFirestoreClient(ctx)
	if err != nil {
		return err									
	}
	a.FireClient = fireClient

	fireAuth, err := config.GetAuthClient(ctx)
	if err != nil {
		return err
	}
	a.FireAuth = fireAuth

	a.ListenPort = envy.Get("REST_PORT", "8080")

	return nil
}

func StartRestServer(a *model.Application) error {
	router := gin.New()

	router.Use(cors.New(md.CORSMiddleware()))

	api.SetRoutes(router, a.FireClient, a.FireAuth)

	err := router.Run(":" + a.ListenPort)
	if err != nil {
		return err
	}

	return nil
}
