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
	sh.Time = time.Now()
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

//搜索求助的数量
func CountSearcSeekhHelp(sh *models.SeekHelp)(count int,err error){
	q:=&models.SeekHelp{
		Type:sh.Type,
		Status:sh.Status,
	}
	err=shdb.Model(q).
		Where("gold >= ? and (title like ? or content like ? )",
		sh.Gold,`%`+sh.Title+`%`,`%`+sh.Title+`%`).Where(q).
		Count(&count).Error
	return
}

//搜索求助
func SearchSeekHelp(sh *models.SeekHelp, limit, page int) (shs *[]models.SeekHelp, err error) {
	shs = new([]models.SeekHelp)
	q:=&models.SeekHelp{
		Type:sh.Type,
		Status:sh.Status,
	}
	err=shdb.Model(q).
		Where("gold >= ? and (title like ? or content like ? )",
		sh.Gold,`%`+sh.Title+`%`,`%`+sh.Title+`%`).Where(q).
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
