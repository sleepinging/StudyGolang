package service

import (
	"net/http"
	"fmt"
	"time"
	"../control"
)

func Login(w http.ResponseWriter, r *http.Request) {
	msg := "登录失败"
	r.ParseForm()
	emails, ok := r.PostForm["Email"]
	if !ok {
		control.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	pwds, ok := r.PostForm["Password"]
	if !ok {
		control.SendRetJson(0, "缺少Password参数", "手动滑稽", w)
		return
	}
	pwd := pwds[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr, "登录：", email, pwd)
	lr := control.CheckLogin(email, pwd)
	if lr == 1 {
		msg = "登录成功"
	}
	control.SendRetJson(lr, msg, "手动滑稽", w)
}
