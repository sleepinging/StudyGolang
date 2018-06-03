package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"twt/mytools"
	"../models"
	"../tools"
)

var (
	dbname = `data\db\login.db` //登录数据库
	db     *gorm.DB             //数据库连接
)

//初始化包
func init() {
	pt, _ := mytools.GetCurrentPath()
	dbname = pt + dbname
	tdb, err := gorm.Open("sqlite3", dbname)
	tools.CheckErr(err)
	db = tdb
	if !db.HasTable(&models.Login{}) {
		db.CreateTable(&models.Login{})
	}
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
