package main

import (
	"awesomeProject/controller"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
)

//note https://www.bilibili.com/video/BV1Fy4y117jJ?p=14 视频课程
func main() {
	cfg, err := tool.ParseConfig("./config/config.json")
	if err != nil {
		//panic(err)
	}
	//new数据库
	tool.GormEngine(cfg)
	engine := gin.Default()
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
