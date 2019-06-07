package Middleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtMiddleware "github.com/iris-contrib/middleware/jwt"
	"iris/app/Models"
	"iris/config"
	"time"
)

/**
 * 验证 jwt
 * @method JwtHandler
 */
func JwtHandler() *jwtMiddleware.Middleware {
	return jwtMiddleware.New(jwtMiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JWT_SECRET), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})

}

type Claims struct {
	Id     int64
	Name   string
	Enable bool
	jwt.StandardClaims
}

// 生成auth-token
func GenerateToken(user *Models.User) (string, error) {
	expireTime := time.Now().Add(60 * time.Second)

	claims := Claims{
		user.Id,
		user.Name,
		user.Enable,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(config.JWT_SECRET))
	return token, err
}