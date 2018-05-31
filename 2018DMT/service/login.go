package service

import (
	"net/http"
	"fmt"
	"time"
	"../control"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Email"][0]
	pwd := r.Form["Password"][0]
	fmt.Println(time.Now(), r.RemoteAddr, "登录：", email, pwd)
	lr := control.CheckLogin(email, pwd)
	msg := "登录失败"
	if lr == 1 {
		msg = "登录成功"
	}
	control.SendRetJson(lr, msg, "", w)
}
