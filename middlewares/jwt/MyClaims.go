package jwt

import "github.com/dgrijalva/jwt-go"

//定义jwt的结构体
type MyClaims struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

