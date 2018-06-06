package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//配置文件的模型
type Config struct {
	//网站根目录
	Wwwroot string `json:"Wwwroot"`

	//端口
	Port int `json:"Port"`

	//数据库信息
	DbInfo dbinfo `json:"DbInfo"`

	//工作类型数据
	JobTypeInfoFile string `json:"JobTypeInfoFile"`
}

type dbinfo struct {
	//邮箱验证的数据库
	EmailVerifyDb string `json:"EmailVerifyDb"`

	//邮箱验证的数据库类型
	EmailVerifyDbType string `json:"EmailVerifyDbType"`

	//登录的数据库
	LoginDb string `json:"LoginDb"`

	//登录的数据库类型
	LoginDbType string `json:"LoginDbType"`

	//工作的数据库
	JobDb string `json:"JobDb"`

	//工作的数据库类型
	JobDbType string `json:"JobDbType"`

	//用户的数据库
	UserDb string `json:"UserDb"`

	//用户的数据库类型
	UserDbType string `json:"UserDbType"`

	//金币的数据库
	GoldDb string `json:"GoldDb"`

	//金币的数据库类型
	GoldDbType string `json:"GoldDbType"`

	//简历的数据库
	CVDb string `json:"CVDb"`

	//简历的数据库类型
	CVDbType string `json:"CVDbType"`

	//博客的数据库
	BlogDb string `json:"BlogDb"`

	//博客的数据库类型
	BlogDbType string `json:"BlogDbType"`

	//帮助的数据库
	HelpDb string `json:"HelpDb"`

	//帮助的数据库类型
	HelpDbType string `json:"HelpDbType"`

	//消息的数据库
	MessageDb string `json:"MessageDb"`

	//消息的数据库类型
	MessageDbType string `json:"MessageDbType"`

	//博客回复的数据库
	BlogReplyDb string `json:"BlogReplyDb"`

	//博客回复的数据库类型
	BlogReplyDbType string `json:"BlogReplyDbType"`

	//帮助回复的数据库
	HelpReplyDb string `json:"HelpReplyDb"`

	//帮助回复的数据库类型
	HelpReplyDbType string `json:"HelpReplyDbType"`
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
