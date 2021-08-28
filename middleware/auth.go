package middleware

import (
	"strings"

	"github.com/Thewalkers2012/BlogBackend/response"
	"github.com/Thewalkers2012/BlogBackend/util/jwt"
	"github.com/gin-gonic/gin"
)

const (
	ContextUserIDKey = "UserID"
)

// JWTAuthorMiddleware 基于 JWT 认证的中间件
func JWTAuthorMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带 Token 的三种方式：1.放在请求头，2.放在请求体，3.放在 URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用 Bearer 开头
		// 这里的具体实现方式要依据你的实际的业务决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ResponseError(ctx, response.CodeNeedLogin)
			ctx.Abort()
			return
		}

		// 按空格进行分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.ResponseError(ctx, response.CodeInvalidToken)
			ctx.Abort()
			return
		}

		// parts[1] 是获取到的 tokenString，我们使用之前定义好的解析 JWT 的函数来解析
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			response.ResponseError(ctx, response.CodeInvalidToken)
			ctx.Abort()
			return
		}

		// 将当前请求的 username 信息保存到请求的上下文中
		ctx.Set(ContextUserIDKey, claims.UserID)
		ctx.Next() // 后续的处理函数可以通过 ctx.Get("username") 来获取当前请求的用户信息
	}
}
