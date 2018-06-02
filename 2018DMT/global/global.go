package global

import (
	"../models"
	"fmt"
	"net/http"
	"twt/mytools"
	"os"
	"io/ioutil"
	"encoding/json"
)

var (
	//当前程序路径
	CurrPath string

	//文件服务器
	RootFileServer http.Handler

	//配置信息
	Config models.Config
)

func init() {
	CurrPath, _ = mytools.GetCurrentPath()
	readconfig()
	RootFileServer = http.FileServer(http.Dir(CurrPath + Config.Wwwroot))
}

func readconfig() {
	cfgpath := CurrPath + `config.json`
	file, err := os.OpenFile(cfgpath, os.O_RDONLY, os.ModeType)
	PanicErr(err, "打开配置文件"+cfgpath)
	bs, err := ioutil.ReadAll(file)
	PanicErr(err, "读取配置文件"+cfgpath)
	err = json.Unmarshal(bs, &Config)
	PanicErr(err, "解析配置文件"+cfgpath)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PanicErr(err error, msg string) {
	if err != nil {
		panic(err)
		fmt.Println(msg, "错误")
	}
}