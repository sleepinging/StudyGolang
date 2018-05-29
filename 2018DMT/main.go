package main

import (
	"twt/mytools"
	"net/http"
	"fmt"
	"./EmailVerify"
)

func startserver() {
	pt, _ := mytools.GetCurrentPath()
	http.Handle("/", http.FileServer(http.Dir(pt+`\wwwroot\`)))
	fmt.Println("服务启动...")
	http.ListenAndServe(":80", nil)
}

func test() {
	email := "237731947@qq.com"
	code := EmailVerify.GenCode()
	fmt.Println(code)
	EmailVerify.UpdateCode(email, code)
	//EmailVerify.SendCode(email)
	fmt.Scanln(&code)
	//fmt.Println(code)
	fmt.Println(EmailVerify.CheckCode(email, code))
}

//程序结束
func Quit() {
	EmailVerify.Cleanup()
}

func main() {
	defer Quit()
	test()
}
