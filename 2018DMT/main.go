package main

import (
	"twt/mytools"
	"net/http"
	"fmt"
	"./control/EmailVerify"
	"./control/unittest"
)

func startserver() {
	pt, _ := mytools.GetCurrentPath()
	http.Handle("/", http.FileServer(http.Dir(pt+`\view\wwwroot\`)))
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
