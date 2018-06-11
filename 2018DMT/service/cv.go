package service

import (
	"net/http"
	"../control/Permission"
	"../dao"
	"../models"
	"encoding/json"
	"fmt"
	"net/url"
)

//添加简历
func AddCV(w http.ResponseWriter, r *http.Request) {
	uid, err := Permission.AddCV(w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	r.ParseForm()
	cvs, err := GetPostString("CV", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cv := new(models.CV)
	err = json.Unmarshal([]byte(cvs), cv)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cv.UserId = uid
	id, err := dao.AddCV(cv)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", fmt.Sprintf("%d", id), w)
}

//修改简历
func UpDataCV(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cid, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = Permission.UpdataCV(cid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cvs, err := GetPostString("CV", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cv := new(models.CV)
	err = json.Unmarshal([]byte(cvs), cv)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = dao.UpdataCV(cid, cv)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", `w(ﾟДﾟ)w`, w)
	return
}

//删除简历
func DeleteCV(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	cid, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = Permission.DeleteCV(cid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = dao.DeleteCV(cid)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", `(⊙o⊙)？`, w)
	return
}

//查看用户简历
func GetUserCV(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := GetGetInt("UserId",queryForm)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = Permission.GetUserCV(uid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cvs, err := dao.GetUserCV(uid)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", cvs, w)
	return
}

//查看简历
func GetCV(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	cid, err := GetGetInt("Id", queryForm)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err = Permission.GetCV(cid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	cv, err := dao.GetCV(cid)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", cv, w)
	return
}
