package model

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type User struct {
	Id       int64  `gorm:primary_key column:"id"`
	Username string `gorm:varchar(11) column:"username"`
	Password string `gorm:varchar(11) column:"password"`
	Status   int64  `gorm:tinyint(2) column:"status"`
	Ctime    LocalTime
}

type LocalTime struct {
	time.Time
}

func (t LocalTime) MarshalJSON() ([]byte, error) {
	//格式化秒
	seconds := t.Unix()
	return []byte(strconv.FormatInt(seconds, 10)), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = LocalTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

func (x User) IsStructureEmpty() bool {
	return reflect.DeepEqual(x, User{})
}
