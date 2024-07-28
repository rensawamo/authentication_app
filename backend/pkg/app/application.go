package app

import (
	"context"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	. "github.com/authentication_app/backend/config"
	"github.com/gobuffalo/envy"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Application struct {
	FireClient  *firestore.Client
	FireAuth    *auth.Client
	RedisClient *redis.Client
	RedisPort   string
	ListenPort  string
	GrpcServer  *grpc.Server
	GrpcPort    string
}

func (a *Application) LoadConfigurations() error {
	ctx := context.Background()

	fireClient, err := GetFirestoreClient(ctx)
	if err != nil {
		return err
	}
	a.FireClient = fireClient

	fireAuth, err := GetAuthClient(ctx)
	if err != nil {
		return err
	}
	a.FireAuth = fireAuth

	// REST now unused
	a.ListenPort = envy.Get("REST_PORT", "8080")
	// gRPC port
	a.GrpcPort = envy.Get("GRPC_PORT", "50051")

	return nil
}
