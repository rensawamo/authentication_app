package api

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go/auth"
	md "github.com/RecepieApp/server/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func SetCache(router *gin.Engine, client *redis.Client) {
	router.POST("/set-cache", func(c *gin.Context) {
		print("Setting cache")
		setUserCache(c, client)
	})

	router.GET("/check-expiration", func(c *gin.Context) {
		checkTokenExpiration(c, client)
	})

}

func CreateCustomToken(ctx *gin.Context, auth *auth.Client) {
	claims := map[string]interface{}{
		"premiumAccount": true,
	}

	token, err := auth.CustomTokenWithClaims(ctx, "some-uid", claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error minting custom token",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

// custom

func SetRoutes(router *gin.Engine, client *firestore.Client, auth *auth.Client, redisClient *redis.Client) {
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Use(func(c *gin.Context) {
    authToken := c.GetHeader("AuthToken")
    md.AuthJWT(auth, authToken)(c)
})

	router.GET("/", func(c *gin.Context) {
		showRecepies(c, client)
	})

	router.POST("/recipe", func(c *gin.Context) {
		addRecepie(c, client)
	})

	router.PATCH("/recipe/:id", func(c *gin.Context) {
		updateRecepie(c, client)
	})

	router.DELETE("/recipes/:id", func(c *gin.Context) {
		deleteRecepie(c, client)
	})

	router.Run()
}
