package Permission

import (
	"../../dao"
	"net/http"
	"../../global"
)

func GetUserGold(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 5 || uid == id {
		f = true
		return
	}
	err = global.NoPermission
	return
}

func SetUserGold(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
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

func AddUserGold(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
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
