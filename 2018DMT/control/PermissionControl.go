package control

import (
	"net/http"
	"time"
	"../dao"
	"../global"
)

//用于发布工作权限检查
func PublishJobPermission(w http.ResponseWriter, r *http.Request) (f bool, err error) {
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

//修改工作权限检查
func UpdateJobPermission(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	f = true
	return
}

//删除工作权限检查
func DeleteJobPermission(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	f = true
	return
}

//查询工作权限检查
func QueryJobPermission(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	f = true
	return
}