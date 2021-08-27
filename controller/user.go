package controller

import (
	"net/http"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": removeTopStruct(errs.Translate(trans)), // 翻译错误
			})
		}
		return
	}

	// 手动对请求参数进行详细的规则校验
	// if len(req.Username) == 0 || len(req.Password) == 0 || len(req.RePassword) == 0 || req.Password != req.RePassword {
	// 	// 请求参数有误
	// 	zap.L().Error("SignUp with invalid param")
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	// 2. 业务处理
	server.SignUp(req)
	// 3. 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
