package controller

import (
	"net/http"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 处理注册请求的函数
func SignUpHandler(ctx *gin.Context) {
	// 1. 获取参数和参数校验
	var req models.ParamSignUp
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// 请求参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}

	// 手动对请求参数进行详细的规则校验

	// 2. 业务处理
	server.SignUp()
	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
