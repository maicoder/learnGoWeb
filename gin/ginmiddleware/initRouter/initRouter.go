package initRouter

import (
	"ginmiddleware/handler"
	"ginmiddleware/middleware"
	"ginmiddleware/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.Logger(), gin.Recovery())

	//router.LoadHTMLGlob(filepath.Join(os.Getenv("TMPL_DIR"),"*"))
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
		userRouter.GET("/profile/", middleware.Auth(), handler.UserProfile)
		userRouter.POST("/update", middleware.Auth(), handler.UpdateUserProfile)
	}

	return router
}
