package route

import(
 "github.com/gin-gonic/gin"
 "awesomeProject/controller"
 "awesomeProject/middlewares/jwt"
 "awesomeProject/tool"
)
func LoadRoute(router *gin.Engine) {
	router.NoRoute(tool.NotFound)
	//不需要权限认证
	v1 := router.Group("/user")
	{
		v1.GET("/login", new(controller.UserController).Login)
	}
	v2 := router.Group("v2").Use(jwt.Auth())
	{
		v2.POST("/modify", new(controller.UserController).Modify)
	}

}
