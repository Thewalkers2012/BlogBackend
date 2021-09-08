package routes

import (
	"net/http"
	"time"

	_ "github.com/Thewalkers2012/BlogBackend/docs" // 千万不要忘了导入把你上一步生成的docs

	"github.com/Thewalkers2012/BlogBackend/controller"
	"github.com/Thewalkers2012/BlogBackend/logger"
	"github.com/Thewalkers2012/BlogBackend/middleware"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Setup(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(time.Second*2, 1))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	v1 := r.Group("/api/v1")

	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	// 登录业务路由
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAuthorMiddleware()) // 应用 JWT 中间件
	// token 测试
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailsHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		// 根据时间或者分数获取帖子列表
		v1.GET("/posts2", controller.GetPostListBySomeHandler)

		// 投票
		v1.POST("/vote", controller.PostVoteHandler)
		// 接口文档路由
		r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
