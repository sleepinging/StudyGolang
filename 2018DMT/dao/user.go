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
	//fmt.Println("用户数据库地址:",logindbname)
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

func GetUserByEmail(email string) (user *models.User, err error) {
	u := &models.User{
		Email: email,
	}
	user = new(models.User)
	userdb.Where(u).First(user)
	if user.Id == 0 {
		err = global.NoSuchUser
		return
	}
	return
}

func GetUserById(id int) (user *models.User, err error) {
	u := &models.User{
		Id: id,
	}
	user = new(models.User)
	userdb.Where(u).First(user)
	if user.Id == 0 {
		err = global.NoSuchUser
		return
	}
	return
}

//获取用户类型
func GetUserType(id int) (tp int) {
	u := &models.User{
		Id: id,
	}
	user := new(models.User)
	userdb.Where(u).First(user)
	tp = user.Type
	if tp == 0 {
		tp = 1
	}
	return
}

//修改用户信息
func UpDateUserInfo(id int, user *models.User) (err error) {
	u := new(models.User)
	userdb.Where(&models.User{Id: id}).First(u)
	if u.Id == 0 {
		err = global.NoSuchUser
		return
	}
	u.CopyUserFromExpt(user, []string{"Id"})
	userdb.Save(u)
	return
}
