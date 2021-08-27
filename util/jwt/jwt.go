package jwt

import (
	"errors"
	"time"

	"github.com/Thewalkers2012/BlogBackend/models"
	"github.com/dgrijalva/jwt-go"
)

var ErrorInvalidToken = errors.New("invalid token")

const TokenExpireDuration = time.Hour * 24 * 7

var jwtKey = []byte("thewalkers")

// MyClaims 自定义声明结构体并内嵌 jwt.StandardClaims
// jwt 包自带的 jwt.StandardClaims 只包含了官方字段
// 我们这里需要额外记录一个 username 字段，所以要自定义结构体
// 如果想要保存更多的信息，都可以添加到这个结构体
type MyCliaims struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// 发放 Token
func ReleaseToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(TokenExpireDuration)
	claims := &MyCliaims{
		UserID:   user.UserID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "thewalkers",
			Subject:   "user token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// 解析 Token
func ParseToken(tokenString string) (*MyCliaims, error) {
	var claims = new(MyCliaims)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return claims, nil
	}

	return nil, ErrorInvalidToken
}
