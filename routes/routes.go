package routes

import (
	"net/http"
	"strings"

	"github.com/Thewalkers2012/BlogBackend/controller"
	"github.com/Thewalkers2012/BlogBackend/logger"
	"github.com/Thewalkers2012/BlogBackend/util/jwt"
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
	r.GET("/ping", JWTAuthorMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户，判断请求头中是否有有效的 JWT token
		c.JSON(http.StatusOK, "ok")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, viper.GetString("app.version"))
	})

	return r
}

// JWTAuthorMiddleware 基于 JWT 认证的中间件
func JWTAuthorMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带 Token 的三种方式：1.放在请求头，2.放在请求体，3.放在 URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
		// 这里的具体实现方式要依据你的实际的业务决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "请求头中auth为空",
			})
			ctx.Abort()
			return
		}

		// 按空格进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg":  "请求头中auth格式有误",
			})
			ctx.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg":  "无效的Token",
			})
			ctx.Abort()
			return
		}

		// 将当前请求的 username 信息保存到请求的上下文中
		ctx.Set("userID", claims.UserID)
		ctx.Next() // 后续的处理函数可以通过 ctx.Get("username") 来获取当前请求的用户信息
	}
}
