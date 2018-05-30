package main

import (
	"net/http"
	"fmt"
	"./control/EmailVerify"
	"./service/onindex"
	"./service/onregister"
)

func startserver() {
	http.HandleFunc("/", onindex.GetIndex)
	http.HandleFunc("/register", onregister.Register)
	http.HandleFunc("/register/sendcode", onregister.SendCode)
	fmt.Println("服务启动...")
	http.ListenAndServe(":80", nil)
}

//程序结束
func Quit() {
	EmailVerify.Cleanup()
}

func main() {
	defer Quit()
	//unittest.Test()
	startserver()
	//test()
}
