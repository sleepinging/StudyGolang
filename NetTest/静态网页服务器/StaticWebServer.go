package main

import (
	"net/http"
	"fmt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name, ok := r.Form["UserName"]
	pwd, ok := r.Form["Password"]
	if !ok {
		fmt.Fprintf(w, "请先登录")
		return
	}
	fmt.Println(name, pwd)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if name[0] == pwd[0] {
		fmt.Fprintf(w, "登录成功")
	} else {
		fmt.Fprintf(w, "登录失败")
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(`E:/临时/wwwroot/`)))
	http.HandleFunc("/login", LoginHandler)
	http.ListenAndServe(":80", nil)
}
