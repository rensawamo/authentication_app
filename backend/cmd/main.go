package main

import (
	"flag"
	"log"
	"net"

	"github.com/authentication_app/backend/pkg/app"
	"github.com/authentication_app/backend/pkg/server"
	"github.com/authentication_app/backend/runtime"
	"github.com/joho/godotenv"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	// .env ファイルをロード
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	flag.Parse()

	a := app.Application{}

	if err := a.LoadConfigurations(); err != nil {
		log.Fatalf("Failed to load configurations: %v", err)
	
	}
	if err := runtime.Start(&a); err != nil {
		log.Fatalf("Failed to start the application: %v", err)
}

	server.SetupGrpcServer(&a)

	lis, err := net.Listen("tcp", ":"+a.GrpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("gRPC server listening on %s", a.GrpcPort)
	if err := a.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

func dieFalse(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
