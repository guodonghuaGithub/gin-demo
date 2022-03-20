package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (user *UserDao) Inserts(name string) {
	fmt.Println(2223333333332)
}
