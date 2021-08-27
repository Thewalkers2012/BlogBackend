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
	req := new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(req); err != nil {
		// 请求参数有误
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断 error 是不是 validator
		msg := getErrorMessage(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}

	// 2. 业务处理
	if err := server.SignUp(req); err != nil {
		zap.L().Info("注册失败", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

// 处理登录的控制器
func LoginHandler(ctx *gin.Context) {
	// 1. 验证参数
	req := new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(req); err != nil {
		// 请求参数有误
		zap.L().Error("Login with valid param", zap.Error(err))
		// 判断 errors 是不是 validator 类型
		msg := getErrorMessage(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": msg,
		})
		return
	}

	// 2. 业务逻辑处理
	if err := server.Login(req); err != nil {
		zap.L().Error("server.Login failed", zap.String("username", req.Username), zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误",
		})
		return
	}

	// 3. 返回相应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "用户登录成功",
	})
}
