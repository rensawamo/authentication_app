package model

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	"google.golang.org/grpc"
)

type Application struct {
	FireClient  *firestore.Client
	FireAuth    *auth.Client
	ListenPort  string
	GrpcServer  *grpc.Server
	GrpcPort    string
}