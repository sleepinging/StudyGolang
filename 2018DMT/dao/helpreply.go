package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
	"../models"
	"../tools"
)

var (
	hrdbname = global.Config.DbInfo.HelpReplyDb     //用户数据库
	hrdbtye  = global.Config.DbInfo.HelpReplyDbType //数据库类型
	hrdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go HelpReplyDbInit()
}

//初始化
func HelpReplyDbInit() {
	hrdbname = global.CurrPath + hrdbname
	//fmt.Println("帮助回复数据库地址:",logindbname)
	tdb, err := gorm.Open(hrdbtye, hrdbname)
	tools.PanicErr(err, "帮助回复数据库初始化")
	hrdb = tdb
	if !hrdb.HasTable(&models.HelpReply{}) {
		hrdb.CreateTable(&models.HelpReply{})
	}
	//fmt.Println("帮助回复数据库初始化完成")
	global.WgDb.Done()
}

//回复帮助
func ReplyHelp(){
	
}
