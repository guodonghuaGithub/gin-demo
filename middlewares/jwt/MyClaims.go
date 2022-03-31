package jwt

import "gihtub.com/dgrijalva/jwt-go"

type MyClaims struct{
 Id int64 `json:"id"`
 jwt.StandardClaims

}
