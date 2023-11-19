package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var (
	// 自定义密钥
	LongTokenSecret = "我是长token的密钥" //salt
	// 失效时间
	ExpireTime_ = 36000 //token expire time
)

// 长token的生成方式应该以唯一id为主,不然改掉名字和密码后新token无法生成

func GenLongToken(claims *JWTClaims) (string, error) {
	//将传入的值先在数据库中查询,如果没有则返回没有该账户,如果有则生成长token,因为在注册时已经先将数据写入表中所以即使更新了也能根据更新的内容查询到唯一id
	fmt.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加入密钥生成最终token
	signedToken, err := token.SignedString([]byte(LongTokenSecret))
	if err != nil {
		return "", errors.New("错误")
	}
	return signedToken, nil
}

func ParseLongToken(LongToken string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(LongToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(LongTokenSecret), nil
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
