package server

import (
	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/util/jwt"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
)

func SignUp(req *models.ParamSignUp) (err error) {
	// 1. 判断用户是否存在
	if err = mysql.CheckUserExist(req.Username); err != nil {
		return err
	}

	// 2. 雪花算法生成 UID
	userID := snowflake.GetID()

	// 3. 构造一个 User 实例
	u := models.User{
		UserID:   userID,
		Username: req.Username,
		Password: req.Password,
	}

	// 4. 保存进数据库
	return mysql.CreateUser(&u)
}

func Login(req *models.ParamLogin) (string, error) {
	// 1. 判断用户是否存在
	user := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := mysql.Login(user); err != nil {
		return "", err
	}
	// 由于传递的是指针，当前的u已经有 user_id
	// 即可生成 jwt token
	return jwt.ReleaseToken(user)
}
