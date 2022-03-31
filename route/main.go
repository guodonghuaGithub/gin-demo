pakeage route

"github.com/gin-gonic/gin"
"awesomeProject/Controller"
"awesomeProject/middleware/jwt"
"awesomeProject/tool"

func LoadRoute(route *gin.Engine){
  route.NoRoute(tool.NotFound)
  v1:= route.Group("/user"){
     v1.GET("/login",new(Controller.UserController).Login)
   
  }
  v2:= route.Group("/v2").Use(jwt.Auth()){
     v2.POST("/modify").new(Controller.UserController).ModifyData)
 }  
  
}
