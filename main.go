package main

import (
	"awesomeProject/controller"
	"awesomeProject/middlewares"
	"awesomeProject/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

//note https://www.bilibili.com/video/BV1Fy4y117jJ?p=14 视频课程
func main() {
	cfg, err := tool.ParseConfig("./config/config.json")
	if err != nil {
		//tool.Failed(ctx *gin.Context, "初始化失败")
		fmt.Println("初始化失败")
		return
	}
	engine := gin.Default()
	//初始化session
	//tool.InitSession(engine)
	//初始化数据库
	tool.GormEngine(cfg)
	//设置全局跨域
	engine.Use(middlewares.CoreHadle())
	//注册路由
	RegisterRout(engine)
	engine.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//注册路由
func RegisterRout(engine *gin.Engine) {
	engine.NoRoute(func(context *gin.Context) { //404 全局
		tool.NotFound(context)
	})
	new(controller.UserController).Router(engine.Group("/user"))
}
