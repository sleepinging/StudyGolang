package dao

import (
	"github.com/jinzhu/gorm"
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

func GoldDbInit() {
	golddbname = global.CurrPath + golddbname
	//fmt.Println("用户数据库地址:",logindbname)
	tdb, err := gorm.Open(golddbtye, golddbname)
	tools.PanicErr(err, "用户数据库初始化")
	golddb = tdb
	if !golddb.HasTable(&models.Gold{}) {
		golddb.CreateTable(&models.Gold{})
	}
	fmt.Println("用户数据库初始化完成")
	global.WgDb.Done()
}
