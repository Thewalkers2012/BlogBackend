package routes

import (
	"net/http"

	"github.com/Thewalkers2012/BlogBackend/controller"
	"github.com/Thewalkers2012/BlogBackend/logger"
	"github.com/Thewalkers2012/BlogBackend/middleware"
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
	// 登录业务路由
	r.POST("/login", controller.LoginHandler)
	// token 测试
	r.GET("/ping", middleware.JWTAuthorMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户，判断请求头中是否有有效的 JWT token
		c.JSON(http.StatusOK, "ok")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	return r
}
