package main

import (
	"flag"
	"log"

	"github.com/authentication_app/backend/model"
	"github.com/authentication_app/backend/pkg/server"
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

	a := model.Application{}
	if err := server.SetUpRestServer(&a); err != nil {
		log.Fatalf("Failed to set up REST server: %v", err)
	}

	server.SetUpGrpcServer(&a)

	go server.StartGrpcServer(&a)
	if err := server.StartRestServer(&a); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
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
