package tool

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Gorm struct {
	*gorm.DB
}

var DbEngine *gorm.DB

func GormEngine(cfg *Config) bool {

	db, err := gorm.Open("mysql", cfg.DataBase.User+":"+cfg.DataBase.PassWord+"@tcp("+cfg.DataBase.host+":"+cfg.DataBase.Port+")/"+cfg.DataBase.DbName+"?"+"charset="+cfg.DataBase.Charts+"&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return false
	}
	DbEngine = db
	return true
}
