package api

import (
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/RecepieApp/server/models"
	"github.com/gin-gonic/gin"
)

func showRecepies(ctx *gin.Context, client *firestore.Client) {
	// コンテキストからuser_idを取得
	userID, exists := ctx.Get("userID")
	if !exists {
			// user_idが見つからない場合、エラーを返す
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
	}
	data, err := models.ReadUserCollection(ctx, client, userID.(string))
	if err != nil {
		ctx.JSON(404, gin.H{
			"Message": "Unable to retrieve data",
		})
	}
	ctx.JSON(200, data)

}

func addRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	msg, status := data.AddRecepie(client)

	ctx.JSON(status, gin.H{
		"message": msg,
	})

}

func updateRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Unable to parse data",
		})
	}

	data.UpdateRecepie(ctx, client)

}

func deleteRecepie(ctx *gin.Context, client *firestore.Client) {
	data := models.Recepie{}

	err := models.UnmarshallRequestBodyToAPIData(ctx.Request.Body, &data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"Message": "Unable to parse data",
		})
	}

	data.DeleteUserRecepie(ctx, client)

}
