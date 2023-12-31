package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	// 自定义密钥
	ShortTokenSecret = "我是短token的密钥" //salt
	// 失效时间
	ExpireTime = 3600 //token expire time
)

func GenShortToken(claims *JWTClaims) (string, error) {
	//在 jwt 生成时使用 jwt.NewWithClaims 方法，需传入 header claim 实例 和 密钥。
	//jwt.SigningMethodHS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加入密钥生成最终token
	signedToken, err := token.SignedString([]byte(ShortTokenSecret))
	if err != nil {
		return "", errors.New("错误")
	}
	return signedToken, nil
}
func ParseShortToken(LongToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(LongToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ShortTokenSecret), nil
	})
	if err != nil {
		return nil, errors.New("错误")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New("错误")
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New("错误")
	}
	return claims, nil
}
