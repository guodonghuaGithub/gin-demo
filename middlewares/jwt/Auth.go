package jwt

import (
	"github.com/gin-gonic/gin"
	"awesomeProject/tool"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			tool.Failed(c, -1, "非法访问")
			c.Abort()
			return
		}
		j := NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			tool.Failed(c, -1, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

