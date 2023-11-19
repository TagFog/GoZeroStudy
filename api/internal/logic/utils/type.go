package utils

import "github.com/dgrijalva/jwt-go"

type JWTClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
	Version  int    `json:"version"`
}
