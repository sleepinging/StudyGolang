package global

import (
	"../models"
	"../tools"
	"net/http"
	"os"
)

var (
	//当前程序所在文件夹路径
	CurrPath string

	//文件文件夹路径
	FileDirPath string

	//文件服务器
	RootFileServer http.Handler

	//配置信息
	Config models.Config

	//加解密的key
	MiKey = `TWT1234567890TWT`
)

func init() {
	CurrPath, _ = tools.GetCurrentPath()
	err := Config.Load(CurrPath + `config.json`)
	tools.PanicErr(err, "加载配置文件")
	Config.Wwwroot = CurrPath + Config.Wwwroot
	RootFileServer = http.FileServer(http.Dir(Config.Wwwroot))
	FileDirPath = Config.Wwwroot + `data\file\`
	if f, _ := tools.PathExists(FileDirPath); !f {
		tools.CheckErr(os.MkdirAll(FileDirPath, os.ModePerm))
	}
}