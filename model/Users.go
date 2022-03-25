package model

import (
	"awesomeProject/param"
	"awesomeProject/tool"
)

type Mode struct {
	Id     int64 `gorm:primary_key column:"id"`
	Status int64 `gorm:"default:0,gt=2"`
	Ctime  tool.LocalTime
}
type Users struct {
	Mode
	param.Users
}

type GetParams struct {
	Users
}
type CreateUser struct {
	GetParams
}

// UpdateUserParams 更新用户
type UpdateUserParams struct {
	Users
	param.UserModify
}

func (CreateUser) TableName() string {
	return "users"
}
func (UpdateUserParams) TableName() string {
	return "users"
}
