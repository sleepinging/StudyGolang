package Permission

import (
	"../../dao"
	"../../global"
	"net/http"
	"time"
)

//用于发布工作权限检查
func PublishJob(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 3 || tp == 4 {
		f = true
		return
	}
	err = global.NoPermission
	return
}

//修改工作权限检查
func UpdateJob(jid int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp < 3 {
		return
	}
	//这个工作是自己发布的
	jb, err := dao.GetJobById(jid)
	if err != nil {
		return
	}
	if jb.Id == uid {
		f = true
		return
	}
	return
}

//删除工作权限检查
func DeleteJob(jid int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp < 3 {
		return
	}
	//管理员
	if tp == 5 {
		f = true
		return
	}
	//这个工作是自己发布的
	jb, err := dao.GetJobById(jid)
	if err != nil {
		return
	}
	if jb.Id == uid {
		f = true
		return
	}
	return
}

//查询工作权限检查
func QueryJob(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	f = true
	return
}

func GetUserIdByCookie(w http.ResponseWriter, r *http.Request) (id int, err error) {
	cookie, err := r.Cookie("user")
	if err != nil {
		if err == http.ErrNoCookie {
			err = global.NotLogin
		}
		return
	}
	id, t, _ := dao.GetUserIdFromCookie(cookie.Value)
	dt := int(time.Now().Sub(t).Seconds())
	if dt > global.MaxCookieTime {
		err = global.LoginCookiePass
		return
	}
	return
}
