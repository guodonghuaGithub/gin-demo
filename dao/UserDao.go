package dao

import (
	"awesomeProject/encryption"
	"awesomeProject/model"
	"github.com/jinzhu/gorm"
)

type UserDao struct {
	*gorm.DB
}

func (user *UserDao) Inserts(users model.GetParams) *model.Users {
	model := model.CreateUser{users}
	user.Create(&model)
	if model.Id == 0 {
		return nil
	}
	return &model.Users
}

func (user *UserDao) Query(username, password string) *model.Users {
	var users model.Users
	user.Where("username = ? and password  = ?", username, password).Find(&users)

	if users.Id <= 0 {
		return nil
	}
	return &users
}

func (user *UserDao) ModifyData(params model.UpdateUserParams) bool {
	var model model.Users
	model.Password = encryption.Md5(params.NewPassword)
	user.Model(&model).Where("id = ?", params.Users.Id).Update(model)
	if model.Id == 0 {
		return false
	}

	return true
}
