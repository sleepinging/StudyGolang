package service

import (
	"net/http"
	"net/url"
	"../models"
	"../control/Permission"
	"../dao"
)

//查看求助
func GetSeekHelp(w http.ResponseWriter, r *http.Request)  {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id,err:=GetGetInt("Id",queryForm)
	uid,err:=Permission.GetSeekHelp(id,w,r)
	_=uid
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	sh,err:=dao.GetSeekHelp(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", sh, w)
	return
}

//发布求助
func PublishSeekHelp(w http.ResponseWriter, r *http.Request)  {
	r.ParseForm()
	shs,err:=GetPostString("SeekHelp",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	sh,err:=models.LoadSeekHelpFromStr(shs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.PublishSeekHelp(sh,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	sid,err:=dao.PublishSeekHelp(sh)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", sid, w)
	return
}

//搜索求助的数量
func CountSearcSeekhHelp(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	shs,err:=GetGetString("SeekHelp",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	sh,err:=models.LoadSeekHelpFromStr(shs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.SearchSeekHelp(sh,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	count,err:=dao.CountSearcSeekhHelp(sh)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", count, w)
}

//搜索求助
func SearchSeekHelp(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit,err:=GetGetInt("Limit",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page,err:=GetGetInt("Page",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	shs,err:=GetGetString("SeekHelp",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	sh,err:=models.LoadSeekHelpFromStr(shs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.SearchSeekHelp(sh,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	res,err:=dao.SearchSeekHelp(sh,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", res, w)
}

//删除求助
func DeleteSeekHelp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	sid, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	uid,err := Permission.DeleteSeekHelp(sid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	_=uid
	err = dao.DeleteSeekHelp(sid)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", `(⊙o⊙)？`, w)
	return
}

//修改求助
func UpdataSeekHelp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	sid, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	uid,err := Permission.UpdataSeekHelp(sid, w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	_=uid
	shs,err:=GetPostString("SeekHelp",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	sh,err:=models.LoadSeekHelpFromStr(shs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err = dao.UpdataSeekHelp(sid,sh)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(0, "成功", `💢`, w)
}
