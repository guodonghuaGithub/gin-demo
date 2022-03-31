package jwt

//awesomeProject/tool"
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
    	"awesomeProject/tool"
	"time"
)

var _cfg = tool.Config{}

const TokenExpireDuration = time.Hour * 2

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(_cfg.Jwt.Search),
	}
}

//生成jtw
func (j *JWT) GenerateToken(Id int64) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		Id, // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    _cfg.Jwt.Issuer,                            // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(j.SigningKey)
}
func (j *JWT) ParseToken(t string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(t, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if v, ok := err.(*jwt.ValidationError); ok {
			if v.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("That's not even a token")
			} else if v.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, errors.New("Token is expired")
			} else if v.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token:")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("Couldn't handle this token:")

	} else {
		return nil, errors.New("Couldn't handle this token:")

	}
}

