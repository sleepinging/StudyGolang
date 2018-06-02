package global

import (
	"../models"
	"../tools"
	"net/http"
	"twt/mytools"
)

var (
	//当前程序路径
	CurrPath string

	//文件服务器
	RootFileServer http.Handler

	//配置信息
	Config models.Config

	//加解密的key
	MiKey = `TWT1234567890TWT`
)

func init() {
	CurrPath, _ = mytools.GetCurrentPath()
	err := Config.Load(CurrPath + `config.json`)
	tools.PanicErr(err, "加载配置文件")
	RootFileServer = http.FileServer(http.Dir(CurrPath + Config.Wwwroot))
}