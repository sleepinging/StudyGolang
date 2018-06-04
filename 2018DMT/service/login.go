package service

import (
	"../control"
	"../models"
	"../dao"
	"../global"
	"fmt"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	msg := "账号与密码不匹配"
	r.ParseForm()
	emails, ok := r.PostForm["Email"]
	if !ok {
		models.SendRetJson(0, "缺少Email参数", "手动滑稽", w)
		return
	}
	email := emails[0]
	pwds, ok := r.PostForm["Password"]
	if !ok {
		models.SendRetJson(0, "缺少Password参数", "手动滑稽", w)
		return
	}
	pwd := pwds[0]
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
		r.RemoteAddr, "登录：", email, pwd)
	lr, user, err := control.CheckLogin(email, pwd)
	if lr == 1 {
		if err != nil {
			models.SendRetJson(lr, "获取用户信息失败", err.Error(), w)
			return
		}
		msg = "登录成功"
		cookie := http.Cookie{
			Name:   "user",
			MaxAge: MaxCookieTime,
			Value:  dao.GenUserCookie(user.ToString()),
		}
		http.SetCookie(w, &cookie)
	}
	models.SendRetJson(lr, msg, "手动滑稽", w)
}

func IsLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		models.SendRetJson(0, "未登录", "", w)
		return
	}
	user, t, err := dao.GetUserinfo(cookie.Value)
	dt := int(time.Now().Sub(t).Seconds())
	if dt > MaxCookieTime {
		models.SendRetJson(1, "登录过期", user, w)
		return
	}
	models.SendRetJson(1, t.Format("2006-01-02/15:04:05"), user, w)
}

//用于服务器内部检查
func Check(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	cookie, err := r.Cookie("user")
	if err != nil {
		return
	}
	user, t, _ := dao.GetUserinfo(cookie.Value)
	dt := int(time.Now().Sub(t).Seconds())
	if dt > MaxCookieTime {
		err = global.LoginCookiePass
		return
	}
	_ = user
	f = true
	return
}
