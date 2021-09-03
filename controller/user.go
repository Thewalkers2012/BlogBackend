package controller

import (
	"errors"
	"fmt"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/response"
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
			response.ResponseError(ctx, response.CodeInvalidParam)
		} else {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	// 2. 业务处理
	if err := server.SignUp(req); err != nil {
		zap.L().Error("注册失败", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			response.ResponseError(ctx, response.CodeUserExist)
		} else {
			response.ResponseError(ctx, response.CodeServerBusy)
		}
		return
	}

	// 3. 返回响应
	response.ResponseSuccess(ctx, nil)
}

// 处理登录的控制器
func LoginHandler(ctx *gin.Context) {
	// 1. 验证参数
	req := new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(req); err != nil {
		// 请求参数有误
		zap.L().Error("Login with valid param", zap.Error(err))
		// 判断 errors 是不是 validator 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(ctx, response.CodeInvalidParam)
		} else {
			response.ResponseErrorWithMsg(ctx, response.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	// 2. 业务逻辑处理
	user, err := server.Login(req)
	if err != nil {
		zap.L().Error("server.Login failed", zap.String("username", req.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			response.ResponseError(ctx, response.CodeUserNotExist)
		} else if errors.Is(err, mysql.ErrorInValidPassword) {
			response.ResponseError(ctx, response.CodeInvalidPassword)
		} else {
			response.ResponseError(ctx, response.CodeServerBusy)
		}
		return
	}

	// 4. 返回相应
	response.ResponseSuccess(ctx, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
