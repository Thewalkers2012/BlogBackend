package controller

import (
	"strconv"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/response"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
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

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(ctx *gin.Context) {
	// 1. 获取参数 （从 URL 中获取帖子的 id）
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid pararm", zap.Error(err))
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}

	// 2. 根据 id 取出帖子中的数据 (查数据库)
	data, err := server.GetPostByID(pid)
	if err != nil {
		zap.L().Error("server.GetPostById(pid) failed", zap.Error(err))
		response.ResponseError(ctx, response.CodeServerBusy)
		return
	}

	// 3. 返回相应
	response.ResponseSuccess(ctx, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(ctx *gin.Context) {
	// 获取参数
	pageNum := ctx.Query("page_num")
	pageSize := ctx.Query("page_size")

	var (
		limit  int64
		offset int64
		err    error
	)
	offset, err = strconv.ParseInt(pageNum, 10, 64)
	if err != nil {
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}

	limit, err = strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}

	// 获取数据
	data, err := server.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("server.GetPostList() failed", zap.Error(err))
		return
	}
	// 返回相应
	response.ResponseSuccess(ctx, data)
}

/**
GetPostListBySomeHandler 根据时间顺序或者是分数大小进行获取帖子的列表
升级版帖子列表接口，按照时间或者按照分数进行排序
1. 获取参数
2. 去 redis 中查询 id 列表
3. 根据 id 去数据库中查询帖子详细信息
*/

// GetPostListBySomeHandler 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetPostListBySomeHandler(ctx *gin.Context) {
	// 获取参数
	// 初始化结构体时，指定初始参数
	req := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: models.OrderTime, // magic string
	}

	if err := ctx.ShouldBindQuery(req); err != nil {
		zap.L().Error("GetPostListBySomeHandler with invalid param", zap.Error(err))
		response.ResponseError(ctx, response.CodeInvalidParam)
		return
	}

	data, err := server.GetPostListNew(req)
	if err != nil {
		zap.L().Error("server.GetPostList() failed", zap.Error(err))
		return
	}
	// 返回相应
	response.ResponseSuccess(ctx, data)
}
