package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context)  {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Hello gin " + strings.ToLower(context.Request.Method) + " Method",
	})
}
