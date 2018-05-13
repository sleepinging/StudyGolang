package main

import (
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name, ok := r.Form["UserName"]
	pwd, ok := r.Form["Password"]
	if !ok {
		//fmt.Fprintf(w, "请先登录<br>")
		http.Redirect(w, r, "/index.html", http.StatusFound)
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
	http.Handle("/", http.FileServer(http.Dir(`E:/developing/wwwroot/`)))
	http.HandleFunc("/login", LoginHandler)
	fmt.Println("服务启动...")
	http.ListenAndServe(":80", nil)
}
