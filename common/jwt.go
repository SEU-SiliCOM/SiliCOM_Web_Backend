package common

import (
	"SilicomAPPv0.3/global"
	"github.com/golang-jwt/jwt"
	"time"
)

var SigningKey = []byte(global.Config.Jwt.Key)

const TokenExpireDuration = time.Hour * 12 //定义JWT过期时间为2小时

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
	C := Claims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), //设置超时时间,输出过期的时间，按照格式
			Issuer:    "admin",                                    //签发人
		},
	}
	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, C)
	//使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(SigningKey)
}

// VerifyToken 验证Token
func VerifyToken(tokenString string) (string, error) {
	MyClaim := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, MyClaim, func(token *jwt.Token) (interface{}, error) {
		return SigningKey, nil
	})
	return MyClaim.Username, err
}
