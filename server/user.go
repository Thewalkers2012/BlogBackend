package server

import (
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
)

func SignUp() {
	// 判断用户是否存在
	mysql.QueryUserByUsername()

	// 生成 UID
	snowflake.GetID()

	// 密码加密

	// 保存进数据库
	mysql.CreateUser()
}
