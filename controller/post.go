package controller

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/response"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(ctx *gin.Context) {
	// 1. 获取参数及参数校验
	req := new(models.Post)
	if err := ctx.ShouldBindJSON(req); err != nil {
		zap.L().Debug("ctx.ShouldBindJSON(req) error", zap.Any("err", err))
		zap.L().Error("create post with invalid param")
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}
	// 从 Context 中获取到当前发请求的用户 id
	userID, err := getCurrentUser(ctx)
	if err != nil {
		response.ResponseError(ctx, response.CodeNeedLogin)
		return
	}
	req.AuthorID = userID

	// 2. 创建帖子
	if err := server.CreatePost(req); err != nil {
		zap.L().Error("server.CreatePost(req) failed", zap.Error(err))
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}

	// 3. 返回相应
	response.ResponseSuccess(ctx, nil)
}
