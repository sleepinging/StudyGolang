package dao

import (
	"../global"
	"../tools"
	"../models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
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
	tools.PanicErr(err, "消息数据库初始化")
	msgdb = tdb
	if !msgdb.HasTable(&models.Message{}) {
		msgdb.CreateTable(&models.Message{})
	}
	//fmt.Println("消息数据库初始化完成")
	global.WgDb.Done()
}

//发送消息
func SendMsg(msg *models.Message) (mid int, err error) {
	msg.Time = time.Now()
	u, err := GetUserById(msg.RecverId)
	if err != nil {
		return
	}
	msg.RecverName = u.Name
	u, err = GetUserById(msg.SenderId)
	if err != nil {
		return
	}
	msg.SenderName = u.Name
	err = msgdb.Create(msg).Error
	if err != nil {
		return
	}
	mid = msg.Id
	return
}

//查看特定消息
func GetMsg(msg *models.Message, limit, page int) (msgs *[]models.Message, err error) {
	msgs = new([]models.Message)
	err = msgdb.Where(msg).Offset((page - 1) * limit).Limit(limit).Find(msgs).Error
	return
}

//标记为已读
func MarkMsgRead(mid int) (err error) {
	msg := new(models.Message)
	err = msgdb.Where(&models.Message{Id: mid}).First(msg).Error
	if err != nil {
		return
	}
	if msg.Id == 0 {
		err = global.NoSuchMsg
		return
	}
	msg.Readed = 1
	err = msgdb.Save(msg).Error
	return
}

//删除消息
func DeleteMsg(mid int) (err error) {
	msg := new(models.Message)
	err = msgdb.Where(&models.Message{Id: mid}).First(msg).Error
	if err != nil {
		return
	}
	if msg.Id == 0 {
		err = global.NoSuchMsg
		return
	}
	err = msgdb.Delete(msg).Error
	return
}

//消息条数
func MsgCount(msg *models.Message) (c int, err error) {
	msgs := new([]models.Message)
	err = msgdb.Where(msg).Find(msgs).Error
	c = len(*msgs)
	return
}

//查找消息
func GetMsgById(mid int) (msg *models.Message, err error) {
	msg = new(models.Message)
	err = msgdb.Where(&models.Message{Id: mid}).First(msg).Error
	return
}
