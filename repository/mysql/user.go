package mysql

import (
	"database/sql"
	"errors"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/util"
)

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) error {
	query := `select count(user_id) from user where username = ?`

	var count int64
	if err := db.Get(&count, query, username); err != nil {
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
	query := `insert into user(user_id, username, password) values(?, ?, ?)`
	_, err = db.Exec(query, user.UserID, user.Username, hashedPassword)
	return
}

// LoginUser 检查用户存在，并且密码相等
func Login(user *models.User) error {
	password := user.Password

	query := `select user_id, username, password from user where username = ?`
	if err := db.Get(user, query, user.Username); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("用户不存在")
		}
		return errors.New("数据库内部错误")
	}

	// 判断密码是否相等
	if err := util.CheckPassword(password, user.Password); err != nil {
		return errors.New("密码错误")
	}

	return nil
}
