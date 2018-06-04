package service

import (
	"net/http"
	"fmt"
	"time"
	"../global"
	"../dao"
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
		cookie := http.Cookie{Name: "user", MaxAge: MaxCookieTime, Value: dao.GenUserCookie(email)}
		http.SetCookie(w, &cookie)
	}
	tools.SendRetJson(lr, msg, "手动滑稽", w)
}

func IsLogin(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("user")
	if err != nil {
		tools.SendRetJson(0, "", "未登录", w)
		return
	}
	user, t, err := dao.GetUserinfo(cookie.Value)
	dt := int(time.Now().Sub(t).Seconds())
	if dt > MaxCookieTime {
		tools.SendRetJson(1, user, "登录过期", w)
		return
	}
	tools.SendRetJson(1, user, t.Format("2006-01-02/15:04:05"), w)
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
