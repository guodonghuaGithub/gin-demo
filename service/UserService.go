package service

import (
	"awesomeProject/dao"
	"awesomeProject/encryption"
	"awesomeProject/model"
	"awesomeProject/tool"
)

type UserService struct {
}

func (user *UserService) Login(LoginParams model.GetParams) *model.Users {
	db := dao.UserDao{tool.DbEngine}
	LoginParams.Password = encryption.Md5(LoginParams.Password)
	find := db.Query(LoginParams.Username, LoginParams.Password)
	if find != nil {
		return find
	}
	//如果不存在直接添加该账号
	result := db.Inserts(LoginParams)
	if result == nil{
		return nil
	}
	return result

}

//更新数据
func (user *UserService) Modify(modify model.UpdateUserParams) bool {
	db := dao.UserDao{tool.DbEngine}
	result := db.ModifyData(modify)
	if result == false {
		return false
	}
	return true
}
