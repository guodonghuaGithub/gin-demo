package model

import "reflect"

type User struct {
	Id       int64  `gorm:primary_key column:"id"`
	Username string `gorm:varchar(11) column:"username"`
	Password string `gorm:varchar(11) column:"password"`
	Status   int64  `gorm:tinyint(2) column:"status"`
}

func (x User) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, User{})
}
