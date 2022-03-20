package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/param"
	"awesomeProject/tool"
)

type UserService struct {
}

func (user *UserService) Login(LoginParams param.UserParam) *model.User {
	db := dao.UserDao{tool.DbEngine}
	params := model.User{}
	params.Username = LoginParams.UserName
	params.Password = LoginParams.PassWord
	//if params.IsStructureEmpty() { // 括号不能去
	//
	//}

	//如果不存在直接添加该账号
	params.Id = db.Inserts(params)
	if params.Id <= 0 {
		return nil
	}
	return &params
}
