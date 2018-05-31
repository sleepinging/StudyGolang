package service

import (
	"../control/EmailVerify"
	"../control"
	"../dao"
	"../models"
	"fmt"
	"net/http"
	"time"
)

func SendCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Email"][0]
	fmt.Println(time.Now(), r.RemoteAddr, email, "请求验证码")
	if dao.ExistLogin(email) {
		control.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	go EmailVerify.SendCode(email)
	control.SendRetJson(1, "验证码已发送", "", w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.Form["Email"][0]
	pwd := r.Form["Password"][0]
	code := r.Form["VerifyCode"][0]
	fmt.Println(time.Now(), r.RemoteAddr, "注册：", email, pwd, code)
	status, msg := EmailVerify.CheckCode(email, code)
	if !status {
		control.SendRetJson(0, msg, "注册失败", w)
		return
	}
	if dao.ExistLogin(email) {
		control.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	dao.AddLogin(&models.Login{Email: email, Password: pwd})
	control.SendRetJson(1, "注册成功", "返回什么好(滑稽)", w)
}
