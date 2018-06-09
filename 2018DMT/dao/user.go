package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	userdbname = global.Config.DbInfo.UserDb     //用户数据库
	userdbtye  = global.Config.DbInfo.UserDbType //数据库类型
	userdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go UserDbInit()
}

func UserDbInit() {
	userdbname = global.CurrPath + userdbname
	//fmt.Println("用户数据库地址:",logindbname)
	tdb, err := gorm.Open(userdbtye, userdbname)
	tools.PanicErr(err, "用户数据库初始化")
	userdb = tdb
	if !userdb.HasTable(&models.User{}) {
		userdb.CreateTable(&models.User{})
	}
	//fmt.Println("用户数据库初始化完成")
	global.WgDb.Done()
}

func AddUser(user *models.User, pwd string) (id int, err error) {
	err = userdb.Create(user).Error
	if err != nil {
		return
	}
	id = user.Id

	//登录表添加
	err = AddLogin(&models.Login{Email: user.Email, Password: pwd})
	if err != nil {
		return
	}

	//金币表添加
	err = SetUserGold(user.Id, 100)

	return
}

func GetUserByEmail(email string) (user *models.User, err error) {
	u := &models.User{
		Email: email,
	}
	user = new(models.User)
	err = userdb.Where(u).First(user).Error
	if err != nil {
		return
	}
	if user.Id == 0 {
		err = global.NoSuchUser
		return
	}
	return
}

func GetUserById(id int) (user *models.User, err error) {
	user = new(models.User)
	err = userdb.First(user,id).Error
	if err != nil {
		err = global.NoSuchUser
		return
	}
	return
}

//获取用户类型
func GetUserType(id int) (tp int) {
	user := new(models.User)
	err := userdb.First(user,id).Error
	if err != nil {
		err = global.NoSuchUser
		return
	}
	tp = user.Type
	if tp == 0 {
		tp = 1
	}
	return
}

//修改用户信息
func UpDateUserInfo(id int, user *models.User) (err error) {
	u := new(models.User)
	err = userdb.Where(&models.User{Id: id}).First(u).Error
	if err != nil {
		return
	}
	if u.Id == 0 {
		err = global.NoSuchUser
		return
	}
	u.CopyUserFromExpt(user, []string{"Id"})
	userdb.Save(u)
	return
}

//删除用户信息
func DeleteUser(id int) (err error) {
	user := new(models.User)
	err = userdb.Where(&models.User{Id: id}).First(user).Error
	if err != nil {
		return
	}
	if user.Id == 0 {
		err = global.NoSuchUser
		return
	}
	userdb.Delete(user)
	return
}
