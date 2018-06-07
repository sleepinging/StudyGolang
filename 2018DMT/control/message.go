package control

import (
	"../models"
	"../dao"
)

//显示发送的消息
func SendedMsg(uid int, limit, page int) (msgs *[]models.Message, err error) {
	msg := &models.Message{
		SenderId: uid,
	}
	msgs, err = dao.GetMsg(msg, limit, page)
	return
}

//显示收到的消息
func ReceivedMsg(uid int, limit, page int) (msgs *[]models.Message, err error) {
	msg := &models.Message{
		RecverId: uid,
	}
	msgs, err = dao.GetMsg(msg, limit, page)
	return
}

//收到的消息条数
func RecvMsgCount(uid int, msg *models.Message) (c int, err error) {
	c, err = dao.MsgCount(&models.Message{
		RecverId: uid,
		Readed:   msg.Readed,
		Type:     msg.Type,
	})
	return
}

//发送的消息条数
func SendMsgCount(uid int, msg *models.Message) (c int, err error) {
	c, err = dao.MsgCount(&models.Message{
		SenderId: uid,
		Readed:   msg.Readed,
		Type:     msg.Type,
	})
	return
}
