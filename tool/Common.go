package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	CODE_200 = 200 //请求成功
	CODE_404 = 404 //找不到页面
	CODE_0   = 0   //请求失败
	CODE_500 = 500 //服务器异常
)

type Common struct {
}

func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    CODE_200,
		"message": "请求成功",
		"data":    v,
	})
	return

}

func Failed(ctx *gin.Context,code int64, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,
		"message": v,
	})
}

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    CODE_404,
		"message": "找不到地址",
	})
	return
}
