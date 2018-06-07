package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
	"../models"
	"../tools"
	"fmt"
)

var (
	golddbname = global.Config.DbInfo.GoldDb     //用户数据库
	golddbtye  = global.Config.DbInfo.GoldDbType //数据库类型
	golddb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go GoldDbInit()
}

//初始化
func GoldDbInit() {
	golddbname = global.CurrPath + golddbname
	//fmt.Println("用户数据库地址:",logindbname)
	tdb, err := gorm.Open(golddbtye, golddbname)
	tools.PanicErr(err, "金币数据库初始化")
	golddb = tdb
	if !golddb.HasTable(&models.Gold{}) {
		golddb.CreateTable(&models.Gold{})
	}
	fmt.Println("金币数据库初始化完成")
	global.WgDb.Done()
}

//获取用户金币
func GetUserGold(uid int) (g int, err error) {
	gr := &models.Gold{UserId: uid}
	fg := new(models.Gold)
	err = golddb.Where(gr).First(fg).Error
	if err != nil {
		return
	}
	if fg.UserId == 0 {
		err = global.NoSuchUser
		return
	}
	g = fg.Gold
	return
}

//修改用户金币,会添加用户
func SetUserGold(uid, g int) (err error) {
	gr := &models.Gold{UserId: uid}
	fg := new(models.Gold)
	golddb.Where(gr).First(fg)
	fg.UserId = uid
	fg.Gold = g
	err = golddb.Save(fg).Error
	return
}

//添加用户金币
func AddUserGold(uid, g int) (err error) {
	gr := &models.Gold{UserId: uid}
	fg := new(models.Gold)
	err = golddb.Where(gr).First(fg).Error
	if err != nil {
		return
	}
	if fg.UserId == 0 {
		err = global.NoSuchUser
		return
	}
	fg.Gold += g
	err = golddb.Save(fg).Error
	return
}