package control

import (
	"net/http"
	"time"
	"../dao"
	"../global"
)

//用于权限检查
func CheckPublishJobPermission(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	cookie, err := r.Cookie("user")
	if err != nil {
		return
	}
	user, t, _ := dao.GetUserinfo(cookie.Value)
	dt := int(time.Now().Sub(t).Seconds())
	if dt > global.MaxCookieTime {
		err = global.LoginCookiePass
		return
	}
	_ = user
	f = true
	return
}
