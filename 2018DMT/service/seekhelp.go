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
	sh,err:=models.LoadSeekHelp(shs)
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
	models.SendRetJson2(1, "成功", sh.Id, w)
	return
}