package initRouter

import (
	"github.com/gin-gonic/gin"
	"ginmysql/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "statics")
	router.StaticFile("/favicon.ico", "favicon.ico")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	return router
}

//func retHelloGinAndMethod(context *gin.Context)  {
//	context.String(http.StatusOK, "hello gin " + strings.ToLower(context.Request.Method) + " method")
//}