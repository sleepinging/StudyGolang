package Permission

import (
	"../../dao"
	"../../global"
	"net/http"
)

//添加博客的权限检查
func PublishBlog(w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp > 1 {
		f = true
		return
	}
	err = global.NoPermission
	return
}

//修改博客的权限检查
func UpDateBlog(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	if uid == id {
		f = true
		return
	}
	return
}

//删除博客的权限检查
func DeleteBlog(id int, w http.ResponseWriter, r *http.Request) (f bool, err error) {
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
