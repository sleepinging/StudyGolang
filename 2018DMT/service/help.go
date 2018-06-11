package service

import (
	"net/http"
	"../models"
	"../dao"
	"../control/Permission"
	"net/url"
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
func GetSeekHelpHelps(w http.ResponseWriter, r *http.Request){
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
	hrs,err:=GetPostString("HelpReply",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	hr,err:=models.LoadHelpReplyFromStr(hrs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.ReplyHelp(hr,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	err=dao.ReplyHelp(hr)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", hr, w)
}

//获取某帮助的回复数量
func CountHelpReply(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id,err:=GetGetInt("Id",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	c,err:=dao.CountHelpReply(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}

//获取某帮助的回复
func GetHelpReply(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id,err:=GetGetInt("Id",queryForm)
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
	hrs,err:=dao.GetHelpReply(id,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功",hrs, w)
}