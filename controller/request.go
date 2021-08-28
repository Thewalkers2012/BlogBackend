package controller

import (
	"errors"

	"github.com/Thewalkers2012/BlogBackend/middleware"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户没有登录")

// getCurrentUser 迅速获得用户的 userid
func getCurrentUser(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(middleware.ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
