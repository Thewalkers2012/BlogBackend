package controller

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/response"
)

// _ResponsePostList 帖子列表接口响应数据
type _ResponsePostList struct {
	Code    response.ResCode        `json:"code"`    // 业务响应状态码
	Message string                  `json:"message"` // 提示信息
	Data    []*models.ApiPostDetail `json:"data"`    // 数据
}
