package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/auth"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// AuthInterceptor returns a new unary server interceptor that performs per-request auth.
func AuthInterceptor(firebaseAuth *auth.Client) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			// Extract the token from the metadata/context.
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
					return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
			}

			values := md["authorization"]
			if len(values) == 0 {
					return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
			}

			// Expect the authorization header to be in the format `Bearer <token>`
			authHeader := values[0]
			if !strings.HasPrefix(authHeader, "Bearer ") {
					return nil, status.Errorf(codes.Unauthenticated, "authorization token is not in the 'Bearer <token>' format")
			}

			idToken := strings.TrimPrefix(authHeader, "Bearer ")

			// Verify the ID token
			token, err := firebaseAuth.VerifyIDToken(ctx, idToken)
			if err != nil {
					return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token: %v", err)
			}

			// Add the token's UID to the context for further use in the handler if needed
			ctx = context.WithValue(ctx, "uid", token.UID)

			// Continue processing the request
			return handler(ctx, req)
	}
}


func AuthJWT(client *auth.Client, authToken string) gin.HandlerFunc {

	return func(c *gin.Context) {
		token := strings.Replace(authToken, "Bearer ", "", 1)

		// Verify the ID token while checking if the token is revoked by passing checkRevoked
		idToken, err := client.VerifyIDTokenAndCheckRevoked(c, token)
		// firebase の user id を取得
		fmt.Println("ID Token", idToken)
		if err != nil {
			log.Printf("Token verification failed: %v", err)

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("userID", idToken.UID)
		c.Next()
	}
}

func CORSMiddleware() cors.Config {
	clientPort := envy.Get("REACT_PORT", "http://localhost:3000")

	corsConfig := cors.Config{
		AllowOrigins:     []string{clientPort},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return corsConfig
}
