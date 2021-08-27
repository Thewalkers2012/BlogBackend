package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
	"code" : 10001 // 程序中的错误码
	"msg": xx     // 提示信息
	"data"： {}    // 数据
*/

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseError(ctx *gin.Context, code ResCode) {
	res := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func ResponseErrorWithMsg(ctx *gin.Context, code ResCode, msg interface{}) {
	res := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}

	ctx.JSON(http.StatusOK, res)
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	res := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}

	ctx.JSON(http.StatusOK, res)
}
