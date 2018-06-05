package service

import (
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
	if !ok || len(emails) == 0 {
		models.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	fmt.Println(time.Now(), r.RemoteAddr, email, "请求验证码")
	if dao.ExistLogin(email) {
		models.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	go EmailVerify.SendCode(email)
	models.SendRetJson(1, "验证码已发送", "", w)
}

func Register(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	emails, ok := r.PostForm["Email"]
	if !ok || len(emails) == 0 {
		models.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	pwds, ok := r.PostForm["Password"]
	if !ok || len(pwds) == 0 {
		models.SendRetJson(0, "缺少Password参数", "手动滑稽", w)
		return
	}
	pwd := pwds[0]
	codes, ok := r.PostForm["VerifyCode"]
	if !ok || len(codes) == 0 {
		models.SendRetJson(0, "缺少VerifyCode参数", "手动滑稽", w)
		return
	}
	code := codes[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr, "注册：", email, pwd, code)
	status, msg := EmailVerify.CheckCode(email, code)
	if !status {
		models.SendRetJson(0, msg, "注册失败", w)
		return
	}
	if dao.ExistLogin(email) {
		models.SendRetJson(0, "该邮箱已被注册", "", w)
		return
	}
	res, err := dao.AddLogin(&models.Login{Email: email, Password: pwd})
	models.SendRetJson(res, "提示", err.Error(), w)
}
