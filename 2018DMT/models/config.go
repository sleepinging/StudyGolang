package models

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//配置文件的模型
type Config struct {
	//网站根目录
	Wwwroot string `json:wwwroot`

	//端口
	Port int `json:port`
}

func (this *Config) Load(filename string) (err error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeType)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("打开配置文件失败:" + err.Error())
		return
	}
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("读取配置文件失败:" + err.Error())
		return
	}
	err = json.Unmarshal(bs, this)
	if err != nil {
		err = fmt.Errorf("解析配置文件失败:" + err.Error())
		return
	}
	return
}

func (this *Config) Save(filename string) {
	bs, _ := json.Marshal(this)
	ioutil.WriteFile(filename, bs, os.ModeAppend)
}
