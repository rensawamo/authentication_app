package main

import (
	"flag"
	"log"
	"net"

	"github.com/authentication_app/backend/pkg/app"
	"github.com/authentication_app/backend/pkg/server"
)

var _ = flag.Bool("debug", false, "Enable Bun Debug log")

func main() {
	flag.Parse()

	a := app.Application{}

	if err := a.LoadConfigurations(); err != nil {
		log.Fatalf("Failed to load configurations: %v", err)
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
