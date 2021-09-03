package controller

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/response"
	"github.com/Thewalkers2012/BlogBackend/server"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 投票

func PostVoteHandler(ctx *gin.Context) {
	// 参数校验
	req := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(req); err != nil {
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			response.ResponseError(ctx, response.CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans)) // 翻译并除掉错误提示中的结构体
		response.ResponseErrorWithMsg(ctx, response.CodeInvalidParam, errData)
		return
	}

	server.PostVote()

	response.ResponseSuccess(ctx, nil)
}
