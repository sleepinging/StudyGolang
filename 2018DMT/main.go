package main

import (
	"./control/EmailVerify"
	"./control/unittest"
	"./global"
	"./service"
	"./tools"
	"fmt"
	"net/http"
	"time"
)

func RegisterAllRouter() {
	http.HandleFunc("/", service.GetIndex)
	http.HandleFunc("/register", service.Register)
	http.HandleFunc("/register/sendcode", service.SendCode)
	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/login/islogin", service.IsLogin)
	http.HandleFunc("/upload/file", service.GetFileUrl)
	http.HandleFunc("/job/publish", service.PublishJob)
	http.HandleFunc("/job", service.ShowJob)
	http.HandleFunc("/job/query", service.QueryJob)
	http.HandleFunc("/job/querycount", service.QueryJobCount)
	http.HandleFunc("/job/updata", service.UpdataJob)
	http.HandleFunc("/job/delete", service.DeleteJob)
}

func startserver(addr string) {
	fmt.Println("添加路由规则中...")
	RegisterAllRouter()
	fmt.Println("服务启动中...")
	fmt.Println("程序根目录:", global.CurrPath)
	fmt.Println("HTTP根目录:", global.Config.Wwwroot)
	fmt.Println("监听端口:", global.Config.Port)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	err := http.ListenAndServe(addr, nil)
	tools.CheckErr(err)
}

//程序结束
func Quit() {
	EmailVerify.Cleanup()
}

func main() {
	defer Quit()
	unittest.Test()
	addr := `:` + fmt.Sprintf("%d", global.Config.Port)
	startserver(addr)
}
