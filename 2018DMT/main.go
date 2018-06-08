package main

import (
	"./control/EmailVerify"
	"./control/unittest"
	"./global"
	"./tools"
	"./service/router"
	"fmt"
	"net/http"
	"time"
)


func init() {
	//在main启动之前肯定都已经add了
	//等待数据库都加载完毕
	global.WgDb.Wait()
	fmt.Println("数据库初始化完成")
}

func startserver(addr string) {
	fmt.Println("添加路由规则中...")
	router.RegisterAllRouter()
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
