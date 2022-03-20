package controller

import (
	"awesomeProject/dao"
	"awesomeProject/service"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (user *UserController) Router(engine *gin.RouterGroup) {
	engine.GET("/login", user.readEndpoint)
}

func (user *UserController) readEndpoint(c *gin.Context) {
	ms := service.UserService{}
	demo := ms.Login("123")
	dao := dao.UserDao{tool.DbEngine}
	dao.Inserts("123")
	c.JSON(200, gin.H{

		"message": "readEndpoint",
		"data":    demo,
	})
	c.Writer.Write([]byte("返回的数据结果为：" + demo))
}
