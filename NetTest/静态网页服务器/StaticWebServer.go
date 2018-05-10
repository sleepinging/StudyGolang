package main

import (
	"net/http"
	"fmt"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "登录成功") //这个写入到w的是输出到客户端的
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(`E:/临时/wwwroot/`)))
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":80", nil)
}
