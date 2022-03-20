package dao

import (
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (user *UserDao) Inserts(users model.User) int64 {
	user.Create(&users)
	return users.Id
}
