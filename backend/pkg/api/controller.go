package api

import (
	"context"
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	md "github.com/authentication_app/backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetRoutes(router *gin.Engine, client *firestore.Client, authClient *auth.Client, ) {
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// If the access token is incorrect, error
	router.Use(func(c *gin.Context) {
		authToken := c.GetHeader("AuthToken")
		if authToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "AuthToken required"})
			c.Abort()
			return
		}

		md.AuthJWT(authClient, authToken)(c)

		c.Next()
	})

	router.POST("/getCustomToken", func(c *gin.Context) {
		authToken := c.GetHeader("AuthToken")
		if authToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "AuthToken required"})
			return
		}

		customToken, err := generateCustomToken(authClient, authToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate custom token"})
			return
		}
	
		c.JSON(http.StatusOK, gin.H{"customToken": customToken})
	})
	router.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.Run()
}

// generateCustomToken はアクセストークンを受け取り、カスタムトークンを生成
func generateCustomToken(authClient *auth.Client, accessToken string) (string, error) {
	ctx := context.Background()
	token, err := authClient.VerifyIDToken(ctx, accessToken)
	if err != nil {
		return "", err
	}

	uid := token.UID
	customToken, err := authClient.CustomToken(ctx, uid)
	if err != nil {
		return "", err
	}

	return customToken, nil
}

