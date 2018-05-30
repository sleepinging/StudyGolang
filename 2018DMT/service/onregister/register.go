package onregister

import (
	"../../control/EmailVerify"
	"../../control"
	"fmt"
	"net/http"
)

func SendCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Eamil"][0]
	fmt.Println("请求验证码：", email)
	code := EmailVerify.GenCode()
	EmailVerify.SendCode(email)
	control.SendRetJson(1, "", code, w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Eamil"][0]
	pwd := r.Form["Password"][0]
	code := r.Form["VerifyCode"][0]
	fmt.Println("注册：", email, pwd, code)
	status, msg := EmailVerify.CheckCode(email, code)
	su := 0
	if status {
		su = 1
	}
	control.SendRetJson(su, msg, "", w)
}
