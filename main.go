package main

import (
	"awesomeProject/controller"
	"awesomeProject/tool"
	"github.com/gin-gonic/gin"
)

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
	new(controller.UserController).Router(engine.Group("/user"))
}
