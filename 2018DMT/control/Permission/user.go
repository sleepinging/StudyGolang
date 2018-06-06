package Permission

import (
	"net/http"
	"../../dao"
	"../../global"
)

//查看用户的权限
func GetUser(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	f = true
	return
}

//添加用户的权限检查
func AddUser(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 5 {
		f = true
		return
	}
	err = global.NoPermission
	return
}

//修改用户的权限检查
func UpDateUser(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	if uid == id {
		f = true
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 5 {
		f = true
		return
	}
	err = global.NoPermission
	return
}

//删除用户的权限检查
func DeleteUser(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 5 {
		f = true
		return
	}
	err = global.NoPermission
	return
}
