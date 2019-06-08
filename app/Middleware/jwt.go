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
	Id        int64
	Name      string
	NickName  string
	Avatar    string
	Phone     string
	RoleId    int64
	Enable    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	jwt.StandardClaims
}

// 生成auth-token
func GenerateToken(user *Models.User) (string, error) {
	expireTime := time.Now().Add(60 * time.Second)

	claims := Claims{
		user.Id,
		user.Name,
		user.NickName,
		user.Avatar,
		user.Phone,
		user.RoleId,
		user.Enable,
		user.CreatedAt,
		user.UpdatedAt,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "iris-casbins-jwt",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString([]byte(config.JWT_SECRET))
	return token, err
}
