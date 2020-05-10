package initRouter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"ginupload/handler"
	"ginupload/utils"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.StaticFile("/favicon.ico", "favicon.ico")
	router.Static("/statics", "statics")
	router.StaticFS("/avatar", http.Dir(utils.RootPath()+"avatar/"))
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
		userRouter.GET("/profile/", handler.UserProfile)
		userRouter.POST("/update", handler.UpdateUserProfile)
	}

	return router
}
