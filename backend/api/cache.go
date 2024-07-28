package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func getUserCache(ctx *gin.Context, client *redis.Client) string {
	userID := ctx.Query("userID")
	authToken, err := models.GetUserCacheToken(ctx, client, userID)
	if err != nil {
		log.Printf("Issues retriving  Cached Token %v", err)
		return ""
	}

	return authToken
}

func setUserCache(ctx *gin.Context, client *redis.Client) {
	var userCache models.UserCache

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &userCache)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to parse data",
		})
		return
	}

	key := fmt.Sprintf("user:%s", userCache.UserID)
	fmt.Println("Key", key)
	
	headerall, notExists := client.HGetAll(ctx, key).Result()
	fmt.Println("headerall", headerall)
	if notExists == nil {
		userCache.SetCachedToken(ctx, client, key)
		return
	}
}


func checkTokenExpiration(ctx *gin.Context, client *redis.Client) {
	//userID := ctx.Query("userID")
	//key := fmt.Sprintf("user:%s", userID)
	//var expired bool
	//
	//expirationTime := client.ExpireTime(ctx, key).Val()
	//currentTime := time.Now()
	//
	//expired := expirationTime
	//ctx.JSON(200, gin.H{
	//	"message": expired,
	//})
}
