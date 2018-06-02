package service

import (
	"net/http"
	"fmt"
	"time"
	"../control"
	"../tools"
)

func Login(w http.ResponseWriter, r *http.Request) {
	msg := "账号与密码不匹配"
	r.ParseForm()
	emails, ok := r.PostForm["Email"]
	if !ok {
		tools.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	pwds, ok := r.PostForm["Password"]
	if !ok {
		tools.SendRetJson(0, "缺少Password参数", "手动滑稽", w)
		return
	}
	pwd := pwds[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr, "登录：", email, pwd)
	lr := control.CheckLogin(email, pwd)
	if lr == 1 {
		msg = "登录成功"
	}
	tools.SendRetJson(lr, msg, "手动滑稽", w)
}
