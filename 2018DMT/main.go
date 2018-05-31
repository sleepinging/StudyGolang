package main

import (
	"net/http"
	"fmt"
	"./control/EmailVerify"
	"./service"
	"./control/unittest"
	"./global"
)

func startserver() {
	http.HandleFunc("/", service.GetIndex)
	http.HandleFunc("/register", service.Register)
	http.HandleFunc("/register/sendcode", service.SendCode)
	http.HandleFunc("/login", service.Login)
	fmt.Println("服务启动...")
	err := http.ListenAndServe(":80", nil)
	global.CheckErr(err)
}

//程序结束
func Quit() {
	EmailVerify.Cleanup()
}

func main() {
	defer Quit()
	go unittest.Test()
	startserver()
	//test()
}
