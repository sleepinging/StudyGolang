package service

import (
	"net/http"
	"../models"
	"../dao"
	"../control/Permission"
)

//发布帮助
func PublishHelp(w http.ResponseWriter, r *http.Request){
	uid,err:=Permission.PublishHelp(w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	r.ParseForm()
	_=uid
	hs,err:=GetPostString("Help",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	help,err:=models.LoadHelpFromStr(hs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err=dao.PublishHelp(help)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, `成功`, help, w)
	return
}

//获取某求助的帮助数量
func GetHelpCount(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	sid,err:=GetPostInt("Id",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	c,err:=dao.GetHelpCount(sid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
	return
}

//查看某求助的帮助
func WatchSeekHelpHelps(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	id,err:=GetPostInt("Id",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit,err:=GetPostInt("Limit",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page,err:=GetPostInt("Page",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	hps,err:=dao.GetSeekHelpHelps(id,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", hps, w)
}

//接受帮助
func AcceptHelp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	hid,err:=GetPostInt("HelpId",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.AcceptHelp(hid,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	err=dao.AcceptHelp(hid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", `滑稽`, w)
}

//回复帮助
func ReplyHelp(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	models.SendRetJson2(0, "滑稽", `最帅的开发者还没写到这里`, w)
}
