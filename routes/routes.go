package routes

import (
	"net/http"

	"github.com/Thewalkers2012/BlogBackend/controller"
	"github.com/Thewalkers2012/BlogBackend/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Setup(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	return r
}
