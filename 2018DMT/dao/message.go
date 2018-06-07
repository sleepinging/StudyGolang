package dao

import (
	"../global"
	"../tools"
	"../models"
	"github.com/jinzhu/gorm"
	"fmt"
)

var (
	msgdbname = global.Config.DbInfo.MessageDb     //用户数据库
	msgdbtye  = global.Config.DbInfo.MessageDbType //数据库类型
	msgdb     *gorm.DB                             //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go MsgDbInit()
}

//初始化
func MsgDbInit() {
	msgdbname = global.CurrPath + msgdbname
	//fmt.Println("用户数据库地址:",logindbname)
	tdb, err := gorm.Open(msgdbtye, msgdbname)
	tools.PanicErr(err, "简历数据库初始化")
	msgdb = tdb
	if !msgdb.HasTable(&models.Message{}) {
		msgdb.CreateTable(&models.Message{})
	}
	fmt.Println("简历数据库初始化完成")
	global.WgDb.Done()
}

//发送消息
func SendMsg(msg *models.Message) (mid int, err error) {
	err = msgdb.Create(msg).Error
	if err != nil {
		return
	}
	mid = msg.Id
	return
}

//查看某用户消息
func GetUserMsg(uid int) (msgs *[]models.Message, err error) {
	return
}

//标记为已读
func MarkMsgRead(mid int) (err error) {
	return
}

//删除消息
func DeleteMsg(mid int) (err error) {
	return
}
