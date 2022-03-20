package controller

import (
	"awesomeProject/param"
	"awesomeProject/service"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user *UserController) Router(engine *gin.RouterGroup) {
	engine.GET("/login", user.Login)
}

func (user *UserController) Login(c *gin.Context) {
	var UserParam param.UserParam                  //定义参数绑定
	err := tool.Decode(c.Request.Body, &UserParam) //接收机参数
	if err != nil {

		tool.Failed(c, "参数解析失败")
		return
	}
	ms := service.UserService{}
	result := ms.Login(UserParam)
	if result == nil {
		tool.Failed(c, result)
		return
	}
	tool.Success(c, result)
}
