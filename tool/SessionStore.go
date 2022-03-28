package tool

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

type SessionStore struct {
}

func InitSession(engine *gin.Engine) {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	fmt.Println(store)
	if err != nil {
		fmt.Println(err.Error())
	}
	engine.Use(sessions.Sessions("mysession", store))
}

//session-set
func SetSession(ctx *gin.Context, key, v interface{}) error {
	sessions := sessions.Default(ctx)
	sessions.Set(key, v)
	return sessions.Save()
}

//session-get
func GetSession(ctx *gin.Context, k interface{}) interface{} {
	sessions := sessions.Default(ctx)
	return sessions.Get(k)
}
