package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName  string   `json:"app_name"`
	AppPort  string   `json:"app_port"`
	AppHost  string   `json:"app_host"`
	AppModel string   `json:"app_model"`
	DataBase DataBase `json:"DataBase"`
	Jwt	Jwt `json:"jwt"`
}

type DataBase struct {
	Driver   string `josn:"driver"`
	User     string `josn:"user"`
	PassWord string `json:"password"`
	Port     string `json:"port"`
	host     string `json:"host"`
	DbName   string `json:"db_name"`
	Charts   string `json:"charts"`
}
type Jwt struct{
  Search string `json:"search"`
  Issuer string `json:"issuer"`

}
var _Cfg *Config = nil

func ParseConfig(path string) (*Config, error) {
	//fmt.Println(path)//
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(2222)
		//panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	decode := json.NewDecoder(reader)
	if err = decode.Decode(&_Cfg); err != nil {
		return nil, err
	}
	return _Cfg, nil
}
