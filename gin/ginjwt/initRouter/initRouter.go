package initRouter

import (
	"ginjwt/handler/user"
	"ginjwt/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	router.GET("/login", user.CreateJwt)
	router.POST("/register", user.Register)

	return router
}
