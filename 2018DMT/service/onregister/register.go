package onregister

import (
	"../../control/EmailVerify"
	"../../control"
	"fmt"
	"net/http"
)

func SendCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Email"][0]
	fmt.Println("请求验证码：", email)
	EmailVerify.SendCode(email)
	control.SendRetJson(1, "", "", w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Email"][0]
	pwd := r.Form["Password"][0]
	code := r.Form["VerifyCode"][0]
	fmt.Println("注册：", email, pwd, code)
	status, msg := EmailVerify.CheckCode(email, code)
	su := 0
	if status {
		su = 1
	}
	control.SendRetJson(su, msg, "注册成功", w)
}
