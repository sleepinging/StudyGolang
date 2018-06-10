package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
	"../models"
	"../tools"
	"time"
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
func ReplyHelp(hr *models.HelpReply)(err error){
	help,err:=GetHelpById(hr.HelpId)
	if err != nil {
		return
	}
	_=help
	u,err:=GetUserById(hr.ReplyerId)
	if err != nil {
		return
	}
	hr.ReplyerName=u.Name
	u,err=GetUserById(hr.AtId)
	if err != nil {
		return
	}
	hr.AtName=u.Name
	hr.Time=time.Now()
	err=hrdb.Create(hr).Error
	if err != nil {
		return
	}
	return
}

//获取某帮助的回复数量
func CountHelpReply(hid int)(count int,err error){
	err=hrdb.Model(&models.HelpReply{}).Where("help_id =?",hid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

//获取某帮助的回复
func GetHelpReply(hid int, limit, page int)(hrs *[]models.HelpReply,err error){
	hrs=new([]models.HelpReply)
	err=hrdb.Model(&models.HelpReply{}).Where("help_id =?",hid).
		Offset((page - 1) * limit).Limit(limit).
		Order("time desc").
		Find(hrs).Error
	return
}

