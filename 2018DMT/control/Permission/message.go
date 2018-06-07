package Permission

import (
	"net/http"
	"../../models"
	"../../dao"
	"../../global"
)

//发送消息
func SendMsg(msg *models.Message, w http.ResponseWriter, r *http.Request) (err error) {
	_, err = dao.GetUserById(msg.RecverId)
	if err != nil || msg.RecverId == 0 {
		return
	}
	uid, err := GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	msg.SenderId = uid
	msg.Readed = -1
	if tp <= 1 {
		err = global.NoPermission
		return
	}
	return
}

//查询收到的消息条数
func GetRecvedMsg(w http.ResponseWriter, r *http.Request) (uid int, err error) {
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

//查询发送的消息条数
func SendMsgCount(w http.ResponseWriter, r *http.Request) (uid int, err error) {
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

//显示收到的消息
func RecvMsgCount(w http.ResponseWriter, r *http.Request) (uid int, err error) {
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

//显示发送的消息
func GetSendedMsg(w http.ResponseWriter, r *http.Request) (uid int, err error) {
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

//显示某条消息
func GetMsg(mid int, w http.ResponseWriter, r *http.Request) (msg *models.Message, err error) {
	uid, err := GetUserIdByCookie(w, r)
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
	msg, err = dao.GetMsgById(mid)
	if err != nil {
		return
	}
	if msg.SenderId == uid || msg.RecverId == uid {
		return
	}
	return
}

//标为已读
func MarkMsgRead(mid int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
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
	msg, err := dao.GetMsgById(mid)
	if msg.RecverId != uid {
		err = global.NotYourMsg
		return
	}
	return
}

//删除消息
func DeleteMsg(mid int, w http.ResponseWriter, r *http.Request) (err error) {
	uid, err := GetUserIdByCookie(w, r)
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
	msg, err := dao.GetMsgById(mid)
	if msg.RecverId != uid {
		err = global.NotYourMsg
		return
	}
	return
}
