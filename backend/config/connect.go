package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func firebaseApp(ctx context.Context) (*firebase.App, error) {
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_ACCOUNT_KEY_LOCATION"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func GetFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := firebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func GetAuthClient(ctx context.Context) (*auth.Client,  error) {
	app, err := firebaseApp(ctx)
	if err != nil {
		return nil, err
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}
	claims := map[string]interface{}{
		"premiumAccount": true,
}

token, err := authClient.CustomTokenWithClaims(ctx, "some-uid", claims)
if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
}

log.Printf("Got custom token: %v\n", token)
	return authClient, nil
}

