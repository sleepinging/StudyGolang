package service

import (
	"../tools"
	"../control/EmailVerify"
	"../dao"
	"../models"
	"fmt"
	"net/http"
	"time"
)

func SendCode(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	emails, ok := r.PostForm["Email"]
	if !ok {
		tools.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	fmt.Println(time.Now(), r.RemoteAddr, email, "请求验证码")
	if dao.ExistLogin(email) {
		tools.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	go EmailVerify.SendCode(email)
	tools.SendRetJson(1, "验证码已发送", "", w)
}

func Register(w http.ResponseWriter, r *http.Request) {
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
	codes, ok := r.PostForm["VerifyCode"]
	if !ok {
		tools.SendRetJson(0, "缺少VerifyCode参数", "手动滑稽", w)
		return
	}
	code := codes[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr, "注册：", email, pwd, code)
	status, msg := EmailVerify.CheckCode(email, code)
	if !status {
		tools.SendRetJson(0, msg, "注册失败", w)
		return
	}
	if dao.ExistLogin(email) {
		tools.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	dao.AddLogin(&models.Login{Email: email, Password: pwd})
	tools.SendRetJson(1, "注册成功", "返回什么好(滑稽)", w)
}
