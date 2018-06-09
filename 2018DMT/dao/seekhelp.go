package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

var (
	shdbname = global.Config.DbInfo.SeekHelpDb     //用户数据库
	shdbtye  = global.Config.DbInfo.SeekHelpDbType //数据库类型
	shdb     *gorm.DB                              //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go SeekHelpDbInit()
}

//初始化
func SeekHelpDbInit() {
	shdbname = global.CurrPath + shdbname
	tdb, err := gorm.Open(shdbtye, shdbname)
	tools.PanicErr(err, "求助数据库初始化")
	shdb = tdb
	if !shdb.HasTable(&models.SeekHelp{}) {
		shdb.CreateTable(&models.SeekHelp{})
	}
	//fmt.Println("求助数据库初始化完成")
	global.WgDb.Done()
}

//发布求助
func PublishSeekHelp(sh *models.SeekHelp) (sid int, err error) {
	sh.Time = time.Now()
	u, err := GetUserById(sh.PublisherId)
	if err != nil {
		return
	}
	sh.PublisherName = u.Name
	sh.PublisherHead = u.Head
	sh.Status = -1
	if sh.Gold > 0 {
		err = AddUserGold(u.Id, -sh.Gold)
		if err != nil {
			return
		}
	}

	err = shdb.Create(sh).Error
	sid = sh.Id
	return
}

//查看求助
func GetSeekHelp(sid int) (sh *models.SeekHelp, err error) {
	sh = new(models.SeekHelp)
	err = shdb.First(sh, sid).Error
	if err != nil {
		err = global.NoSuchSeekHelp
		return
	}
	return
}

//搜索求助
func SearchSeekHelp(sh *models.SeekHelp, limit, page int) (count int, shs *[]models.SeekHelp, err error) {
	shs = new([]models.SeekHelp)
	//err = shdb.
	//	Where("type = ?", sh.Type).
	//	Where("gold > ? and time > ?", sh.Gold, sh.Time).
	//	Find(shs).
	//	Where("title like ?", `%`+sh.Title+`%`).
	//	Or("content like ?", `%`+sh.Title+`%`).
	//	Offset((page - 1) * limit).Limit(limit).
	//	Find(shs).
	//	Count(&count).
	//	Error
	err = shdb.
		Where("type = ? and status = ? and gold >= ? and time >= ? and (title like ? or content like ? )",
			sh.Type,sh.Status, sh.Gold, sh.Time,`%`+sh.Title+`%`,`%`+sh.Title+`%`).
		Offset((page - 1) * limit).Limit(limit).
		Order("time desc").
		Find(shs).
		Error
	return
}

//删除求助
func DeleteSeekHelp(sid int) (err error) {
	if sid == 0 {
		err = global.NoSuchSeekHelp
		return
	}
	err = shdb.Delete(&models.SeekHelp{Id: sid}).Error
	if err != nil {
		err = global.NoSuchSeekHelp
		return
	}
	return
}

//修改求助
func UpdataSeekHelp(sid int, nsh *models.SeekHelp) (err error) {
	sh := new(models.SeekHelp)
	err = shdb.Where(&models.SeekHelp{Id: sid}).First(sh).Error
	if err != nil {
		err = global.NoSuchSeekHelp
		return
	}
	dg := sh.Gold - nsh.Gold //减少的金币
	if dg != 0 {
		err = AddUserGold(sh.PublisherId, dg)
		if err != nil {
			return
		}
	}
	sh.CopySeekHelpFromExpt(nsh, []string{"Id", ""})
	sh.Time = time.Now()
	err = shdb.Save(sh).Error
	if err != nil {
		err = global.NoSuchSeekHelp
		return
	}
	return
}

//解决求助
func SolveSeekHelp(sid int)(err error){
	sh,err:=GetSeekHelp(sid)
	if err != nil {
		return
	}
	err=shdb.Model(sh).Update("status",1).Error
	return
}
