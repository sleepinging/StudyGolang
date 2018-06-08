package Permission

import (
	"../../dao"
	"../../global"
	"../../models"
	"net/http"
)

//发布求助
func PublishSeekHelp(sh *models.SeekHelp, w http.ResponseWriter, r *http.Request) (uid int, err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp <= 1 {
		err = global.NoPermission
		return
	}
	sh.Id = 0
	sh.Status = -1
	if sh.Gold < 0 {
		err = global.GoldMustPositive
		return
	}
	return
}

//查看求助
func GetSeekHelp(sid int, w http.ResponseWriter, r *http.Request) (uid int, err error) {
	uid, err = GetUserIdByCookie(w, r)
	err = nil
	return
}

//搜索求助
func SearchSeekHelp(sh *models.SeekHelp, w http.ResponseWriter, r *http.Request) (uid int, err error) {
	uid, err = GetUserIdByCookie(w, r)
	err = nil
	return
}

//删除求助
func DeleteSeekHelp(sid int, w http.ResponseWriter, r *http.Request) (uid int, err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	sh,err:=dao.GetSeekHelp(sid)
	if err != nil {
		return
	}
	if uid != sh.PublisherId {
		err = global.NoPermission
		return
	}
	return
}

//修改求助
func UpdataSeekHelp(sid int, w http.ResponseWriter, r *http.Request) (uid int, err error) {
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp >= 5 {
		return
	}
	sh,err:=dao.GetSeekHelp(sid)
	if err != nil {
		return
	}
	if uid != sh.PublisherId {
		err = global.NoPermission
		return
	}
	//TODO 如果有人帮助了
	//dao.getHelpBySeekHelp(sh.Id)
	return
}
