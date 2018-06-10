package Permission

import (
	"net/http"
	"../../dao"
	"../../models"
	"../../global"
)

//发布帮助
func PublishHelp(w http.ResponseWriter, r *http.Request)(uid int,err error){
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	if tp <= 1 {
		err = global.NoPermission
		return
	}
	return
}

//接受帮助
func AcceptHelp(hid int,w http.ResponseWriter, r *http.Request)(uid int,err error){
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	help,err:=dao.GetHelpById(hid)
	if err != nil {
		return
	}
	sh,err:=dao.GetSeekHelp(help.SeekHelpId)
	if uid!=sh.Id{
		err=global.NotYourSeekHelp
		return
	}
	return
}

//回复帮助
func ReplyHelp(hr *models.HelpReply,w http.ResponseWriter, r *http.Request)(uid int,err error){
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	hr.ReplyerId=uid
	if tp<=0{
		err=global.NoPermission
		return
	}
	return
}
