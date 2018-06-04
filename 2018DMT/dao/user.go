package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

var (
	userdbname = global.Config.DbInfo.UserDb     //用户数据库
	userdbtye  = global.Config.DbInfo.UserDbType //数据库类型
	userdb     *gorm.DB                          //数据库连接
)

func init() {
	userdbname = global.CurrPath + userdbname
	tdb, err := gorm.Open(userdbtye, userdbname)
	tools.PanicErr(err, "用户数据库初始化")
	userdb = tdb
	if !userdb.HasTable(&models.User{}) {
		userdb.CreateTable(&models.User{})
	}
	fmt.Println("用户数据库初始化完成")
}

func AddUser(user *models.User) (id int, err error) {
	userdb.Create(user)
	sql := `SELECT last_insert_rowid() as id;`
	var lid tid
	userdb.Raw(sql).Select("id").Scan(&lid)
	id = lid.Id
	return
}
