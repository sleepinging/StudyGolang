package dao

import (
	"github.com/jinzhu/gorm"
	"../models"
	"../tools"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
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
func PublishSeekHelp(sh *models.SeekHelp)(sid int,err error){
	sh.Time=time.Now()
	u, err := GetUserById(sh.PublisherId)
	if err != nil {
		return
	}
	sh.PublisherName=u.Name
	sh.PublisherHead=u.Head
	if sh.Gold>0{
		err=AddUserGold(u.Id,-sh.Gold)
		if err != nil {
			return
		}
	}

	err=shdb.Create(sh).Error
	sid=sh.Id
	return
}

//查看求助
func GetSeekHelp(sid int)(sh *models.SeekHelp,err error){
	sh=new(models.SeekHelp)
	err=shdb.First(sh,sid).Error
	if err != nil {
		err=global.NoSuchSeekHelp
		return
	}
	return
}
