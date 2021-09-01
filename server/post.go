package server

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
)

func CreatePost(req *models.Post) (err error) {
	// 1. 生成 post ID
	req.ID = snowflake.GetID()
	// 2. 保存到数据库
	// 3. 返回
	return mysql.CreatePost(req)
}
