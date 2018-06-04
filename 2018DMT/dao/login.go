package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"twt/mytools"
	"fmt"
)

var (
	dbname = global.Config.DbInfo.LoginDb           //登录数据库
	dbtye  = global.Config.DbInfo.EmailVerifyDbType //数据库类型
	db     *gorm.DB                                 //数据库连接
)

//初始化包
func init() {
	pt, _ := mytools.GetCurrentPath()
	dbname = pt + dbname
	tdb, err := gorm.Open(dbtye, dbname)
	tools.CheckErr(err)
	db = tdb
	if !db.HasTable(&models.Login{}) {
		db.CreateTable(&models.Login{})
	}
	fmt.Println("登录数据库初始化完成")
}

func CheckLogin(login *models.Login) (res int) {
	lg := models.Login{}
	db.Where(login).First(&lg)
	if lg.Email != "" {
		res = 1
	}
	return
}

func ExistLogin(email string) (res bool) {
	lg := models.Login{}
	db.Where(models.Login{Email: email}).First(&lg)
	if lg.Email != "" {
		res = true
	}
	return
}

func AddLogin(login *models.Login) (res int) {
	db.Save(&login)
	return
}
