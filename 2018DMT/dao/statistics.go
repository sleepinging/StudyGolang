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
	sitsdbname = global.Config.DbInfo.StatisticsDb     //用户数据库
	sitsdbtye  = global.Config.DbInfo.StatisticsDbType //数据库类型
	sitsdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go StatisticsDbDbInit()
}

//初始化
func StatisticsDbDbInit() {
	sitsdbname = global.CurrPath + sitsdbname
	//fmt.Println("博客数据库地址:",logindbname)
	tdb, err := gorm.Open(sitsdbtye, sitsdbname)
	tools.PanicErr(err, "统计数据库初始化")
	sitsdb = tdb
	if !sitsdb.HasTable(&models.RegisterRecord{}) {
		sitsdb.CreateTable(&models.RegisterRecord{})
	}
	if !sitsdb.HasTable(&models.VisitRecord{}) {
		sitsdb.CreateTable(&models.VisitRecord{})
	}
	if !sitsdb.HasTable(&models.LoginRecord{}) {
		sitsdb.CreateTable(&models.LoginRecord{})
	}
	//fmt.Println("统计数据库初始化完成")
	global.WgDb.Done()
}

//添加访问记录
func AddVisitRecord(t time.Time,ip ,page string)(err error){
	err=sitsdb.Create(&models.VisitRecord{Time:t,IP:ip,Page:page}).Error
	return
}

//添加注册记录
func AddRegisterRecord(uid int,t time.Time,ip string)(err error){
	err=sitsdb.Create(&models.RegisterRecord{UserId:uid,Time:t,IP:ip}).Error
	return
}

//添加登录记录
func AddLoginRecord(uid,su int,t time.Time,ip string)(err error){
	err=sitsdb.Create(&models.LoginRecord{UserId:uid,Time:t,IP:ip,Su:su}).Error
	return
}

//查询访问记录数
func VisitRecordCount(d int)(c int,err error){
	qt:=time.Now().AddDate(0,0,-d)
	err=sitsdb.Model(&models.VisitRecord{}).Where("time >= ?",qt).Count(&c).Error
	return
}

//查询注册记录数
func RegisterRecordCount(d int)(c int,err error){
	qt:=time.Now().AddDate(0,0,-d)
	err=sitsdb.Model(&models.RegisterRecord{}).Where("time >= ?",qt).Count(&c).Error
	return
}

//查询登录记录数
func LoginRecordCount(d int)(c int,err error){
	qt:=time.Now().AddDate(0,0,-d)
	err=sitsdb.Model(&models.LoginRecord{}).Where("time >= ?",qt).Count(&c).Error
	return
}
