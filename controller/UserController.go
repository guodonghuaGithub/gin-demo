package controller

import (
	"awesomeProject/middlewares/jwt"
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserController struct {
}

func (user *UserController) Router(engine *gin.RouterGroup) {
	engine.GET("/login", user.Login)
	engine.POST("/modify", user.Modify)
}

//le%NU*lx5(kh
func (user *UserController) Login(c *gin.Context) {
	var UserParam model.GetParams //定义参数绑定
	err := c.BindJSON(&UserParam) //接收机参数
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
	jwt := jwt.jWT{}
	tokenString,err := jwt.GenerateToken(result.ID)
	if tokenString == ""{
	   tool.Failed(c,-1,"token创建失败")
	   return	
	}
	tool.Success(c, result)
}

func (user *UserController) Modify(ctx *gin.Context) {
	var UserParams model.UpdateUserParams
	err := ctx.ShouldBindBodyWith(&UserParams, binding.JSON)
	if err != nil {
		fmt.Println(err.Error())
		tool.Failed(ctx, "参数解析失败")
		return
	}
	service := service.UserService{}
	result := service.Modify(UserParams)
	if result == false {
		tool.Failed(ctx, "更新失败")
		return
	}
	tool.Success(ctx, "更新成功")
}

