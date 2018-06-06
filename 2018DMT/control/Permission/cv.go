package Permission

import (
	"../../dao"
	"net/http"
	"../../global"
)

//获取用户的简历
func GetUserCV(id int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 3 || uid == id {
		return
	}
	err = global.NoPermission
	return
}

//获取简历
func GetCV(id int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	cv, err := dao.GetCV(id)
	cuid := cv.UserId
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if uid == cuid || tp >= 3 {
		return
	}
	err = global.NoPermission
	return
}

//修改简历
func UpdateCV(id int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	cv, err := dao.GetCV(id)
	cuid := cv.UserId
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if uid == cuid || tp == 5 {
		return
	}
	err = global.NoPermission
	return
}

//删除简历
func DeleteCV(id int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	cv, err := dao.GetCV(id)
	cuid := cv.UserId
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if uid == cuid || tp == 5 {
		return
	}
	err = global.NoPermission
	return
}

//添加简历
func AddCV(w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp == 2 || tp == 5 {
		return
	}
	err = global.NoPermission
	return
}
