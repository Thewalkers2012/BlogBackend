package mysql

import (
	"errors"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/util"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) error {
	sql := `select count(user_id) from user where username = ?`

	var count int64
	if err := db.Get(&count, sql, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}

	return nil
}

// CreateUser 向数据库中插入一条新的用户记录
func CreateUser(user *models.User) (err error) {
	// 对密码进行加密
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	// 执行 sql 语句
	sql := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(sql, user.UserID, user.Username, hashedPassword)
	return
}
