package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Core struct {
}

func CoreHadle() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKey []string
		for key, _ := range context.Request.Header {
			headerKey = append(headerKey, key)
		}
		headerString := strings.Join(headerKey, ",")
		if headerString != "" {
			fmt.Sprintf("ccess-Control-Allow-Origin,Access-Control-Allow-Headers,%s", headerString)
		} else {
			headerString = "ccess-Control-Allow-Origin,Access-Control-Allow-Headers"
		}
		if origin != "" {
			context.Writer.Header().Set("ccess-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			context.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Option Request")
		}
		context.Next()
	}
}
