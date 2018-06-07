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
	http.HandleFunc("/login/logout", service.Logout)

	http.HandleFunc("/upload/file", service.GetFileUrl)

	http.HandleFunc("/job/publish", service.PublishJob)
	http.HandleFunc("/job", service.ShowJob)
	http.HandleFunc("/job/query", service.QueryJob)
	http.HandleFunc("/job/querycount", service.QueryJobCount)
	http.HandleFunc("/job/updata", service.UpdataJob)
	http.HandleFunc("/job/delete", service.DeleteJob)

	http.HandleFunc("/user", service.GetUser)
	http.HandleFunc("/user/add", service.AddUser)
	http.HandleFunc("/user/updata", service.UpdateUser)
	http.HandleFunc("/user/delete", service.DeleteUser)

	http.HandleFunc("/gold/show", service.GetUserGold)
	http.HandleFunc("/gold/set", service.SetUserGold)

	http.HandleFunc("/cv/user", service.GetUserCV)
	http.HandleFunc("/cv", service.GetCV)
	http.HandleFunc("/cv/update", service.UpDateCV)
	http.HandleFunc("/cv/delete", service.DeleteCV)
	http.HandleFunc("/cv/add", service.AddCV)

	http.HandleFunc("/message", service.GetMsg)
	http.HandleFunc("/message/send", service.SendMsg)
	http.HandleFunc("/message/recved/count", service.RecvMsgCount)
	http.HandleFunc("/message/sended/count", service.SendMsgCount)
	http.HandleFunc("/message/sended", service.GetSendedMsg)
	http.HandleFunc("/message/recved", service.GetRecvedMsg)
	http.HandleFunc("/message/mark", service.MarkMsgRead)
	http.HandleFunc("/message/delete", service.DeleteMsg)
}

func init() {
	//在main启动之前肯定都已经add了
	//等待数据库都加载完毕
	global.WgDb.Wait()
	fmt.Println("数据库初始化完成")
}

func startserver(addr string) {
	fmt.Println("添加路由规则中...")
	RegisterAllRouter()
	fmt.Println("服务启动中...")
	//fmt.Println("程序根目录:", global.CurrPath)
	fmt.Println("HTTP根目录:", global.Config.Wwwroot)
	fmt.Println("监听端口:", global.Config.Port)
	fmt.Println("服务启动时间:", time.Now().Format("2006-01-02 15:04:05"))
	err := http.ListenAndServe(addr, nil)
	tools.ShowErr(err)
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
