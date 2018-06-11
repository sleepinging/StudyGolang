package service

import (
	"net/http"
	"../models"
	"../control"
	"../dao"
	"../control/Permission"
	"encoding/json"
	"net/url"
)

//发送消息
func SendMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	msgs, err := GetPostString("Message", w, r)
	if err != nil {
		models.SendRetJson2(0, "发送失败", err.Error(), w)
		return
	}
	msg := new(models.Message)
	err = json.Unmarshal([]byte(msgs), msg)
	if err != nil {
		models.SendRetJson2(0, "发送失败", err.Error(), w)
		return
	}
	err = Permission.SendMsg(msg, w, r)
	if err != nil {
		models.SendRetJson2(0, "发送失败", err.Error(), w)
		return
	}
	mid, err := dao.SendMsg(msg)
	if err != nil {
		models.SendRetJson2(0, "发送失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "发送成功", mid, w)
	return
}

//查询收到的消息
func GetRecvedMsg(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := Permission.GetRecvedMsg(w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit, err := GetGetInt("Limit", queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page, err := GetGetInt("Page",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	msgs, err := control.ReceivedMsg(uid, limit, page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", msgs, w)
	return
}

//查询发送的消息条数
func SendMsgCount(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := Permission.SendMsgCount(w, r)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	msgs, err := GetGetString("Message", queryForm)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	msg := new(models.Message)
	err = json.Unmarshal([]byte(msgs), msg)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	c, err := control.SendMsgCount(uid, msg)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}

//显示收到的消息条数
func RecvMsgCount(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := Permission.RecvMsgCount(w, r)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	msgs, err := GetGetString("Message", queryForm)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	msg := new(models.Message)
	err = json.Unmarshal([]byte(msgs), msg)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	c, err := control.RecvMsgCount(uid, msg)
	if err != nil {
		models.SendRetJson2(0, "查询失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}

//显示发送的消息
func GetSendedMsg(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid, err := Permission.GetSendedMsg(w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit, err := GetGetInt("Limit", queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page, err := GetGetInt("Page", queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	msgs, err := control.SendedMsg(uid, limit, page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", msgs, w)
	return
}

//显示某条消息
func GetMsg(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id, err := GetGetInt("Id", queryForm)
	msg, err := Permission.GetMsg(id, w, r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", msg, w)
}

//标为已读
func MarkMsgRead(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "标记失败", err.Error(), w)
		return
	}
	err = Permission.MarkMsgRead(id, w, r)
	if err != nil {
		models.SendRetJson2(0, "标记失败", err.Error(), w)
		return
	}
	err = dao.MarkMsgRead(id)
	if err != nil {
		models.SendRetJson2(0, "标记失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", "😀", w)
}

//删除消息
func DeleteMsg(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := GetPostInt("Id", w, r)
	if err != nil {
		models.SendRetJson2(0, "删除失败", err.Error(), w)
		return
	}
	err = Permission.DeleteMsg(id, w, r)
	if err != nil {
		models.SendRetJson2(0, "删除失败", err.Error(), w)
		return
	}
	err = dao.DeleteMsg(id)
	if err != nil {
		models.SendRetJson2(0, "删除失败", err.Error(), w)
		return
	}
	models.SendRetJson2(0, "删除成功", "❤", w)
}
