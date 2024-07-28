package runtime

import (
	md "github.com/authentication_app/backend/middleware"
	"github.com/authentication_app/backend/pkg/api"
	"github.com/authentication_app/backend/pkg/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(a *app.Application) error {
	router := gin.New()

	router.Use(cors.New(md.CORSMiddleware()))


	api.SetRoutes(router, a.FireClient, a.FireAuth, a.RedisClient)

	err := router.Run(":" + a.ListenPort)
	if err != nil {
		return err
	}

	return nil
}
