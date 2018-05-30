package main

import (
	"net/http"
	"fmt"
	"./control/EmailVerify"
	"./control/unittest"
	"./service/onindex"
)

func startserver() {
	http.HandleFunc("/", onindex.GetIndex)
	fmt.Println("服务启动...")
	http.ListenAndServe(":80", nil)
}

//程序结束
func Quit() {
	EmailVerify.Cleanup()
}

func main() {
	defer Quit()
	unittest.Test()
	startserver()
	//test()
}
