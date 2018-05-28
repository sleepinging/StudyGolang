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
	fmt.Scanln(&code)
	//fmt.Println(code)
	fmt.Println(EmailVerify.CheckCode(email, code))
}

func main() {
	test()
}
