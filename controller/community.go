package controller

import (
	"strconv"

	"github.com/Thewalkers2012/BlogBackend/response"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ---- 根社区相关的 ----

func CommunityHandler(ctx *gin.Context) {
	// 查询所有的社区 (community_id, community_name) 以列表的形式返回
	data, err := server.GetCommunityList()
	if err != nil {
		zap.L().Error("server.GetCommunityList() failed", zap.Error(err))
		response.ResponseError(ctx, response.CodeServerBusy) // 不轻易把服务端的报错暴露给外面
		return
	}
	response.ResponseSuccess(ctx, data)
}

// CommunityDetailsHandler 社区分类详情
func CommunityDetailsHandler(ctx *gin.Context) {
	// 1. 获取社区 id
	communityID := ctx.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}

	data, err := server.GetCommunityDetails(id)
	if err != nil {
		zap.L().Error("server.GetCommunityDetails() failed", zap.Error(err))
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}

	response.ResponseSuccess(ctx, data)
}
