package Permission

import (
	"net/http"
	"../../dao"
	"../../global"
	"../../models"
)

//发布求助
func PublishSeekHelp(sh *models.SeekHelp,w http.ResponseWriter, r *http.Request) (uid int, err error){
	uid, err = GetUserIdByCookie(w, r)
	if err != nil {
		return
	}
	tp := dao.GetUserType(uid)
	if tp <= 1 {
		err = global.NoPermission
		return
	}
	sh.Id=0
	if sh.Gold<0{
		err=global.GoldMustPositive
		return
	}
	return
}

//查看求助
func GetSeekHelp(sid int,w http.ResponseWriter, r *http.Request) (uid int, err error){
	uid, err = GetUserIdByCookie(w, r)
	err=nil
	return
}

//搜索求助
func SearchSeekHelp(sh *models.SeekHelp,w http.ResponseWriter, r *http.Request)(uid int,err error){
	return
}

//删除求助

//修改求助
