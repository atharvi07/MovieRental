package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoute(engine *gin.Engine) {
	engine.GET("/helloWorld", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})
}
