package service

import (
	"net/http"
	"../control/Permission"
	"../models"
	"../dao"
	"net/url"
	"strconv"
	"encoding/json"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	f, err := Permission.GetUser(w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	if len(queryForm["Id"]) == 0 {
		models.SendRetJson2(0, "失败", "缺少Id参数", w)
		return
	}
	id, err := strconv.ParseInt(queryForm["Id"][0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "id参数需要为整数", "手动滑稽", w)
		return
	}
	user, err := dao.GetUserById(int(id))
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(0, "成功", user, w)
}

//添加用户
func AddUser(w http.ResponseWriter, r *http.Request) {
	f, err := Permission.AddUser(w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
}

//修改用户
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	users, ok := r.PostForm["User"]
	if !ok || len(users) == 0 {
		models.SendRetJson2(0, "手动滑稽", "缺少User参数", w)
		return
	}
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson2(0, "手动滑稽", "缺少Id参数", w)
		return
	}
	user := &models.User{}
	err := json.Unmarshal([]byte(users[0]), user)
	if err != nil {
		models.SendRetJson2(0, "User格式错误", err.Error(), w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	f, err := Permission.UpDateUser(int(id), w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err = dao.UpDateUserInfo(int(id), user)
	if err != nil {
		models.SendRetJson2(0, "修改失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "修改成功", ids[0], w)
	return
}

//删除用户
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	f, err := Permission.DeleteUser(w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	r.ParseForm()
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson2(0, "缺少id参数", "手动滑稽", w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	err = dao.DeleteUser(int(id))
	if err != nil {
		models.SendRetJson2(0, "删除失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "删除成功", ids[0], w)
	return
}
