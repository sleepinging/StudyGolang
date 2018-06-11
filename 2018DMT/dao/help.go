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
	helpdbname = global.Config.DbInfo.HelpDb     //用户数据库
	helpdbtye  = global.Config.DbInfo.HelpDbType //数据库类型
	helpdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go HelpDbInit()
}

//初始化
func HelpDbInit() {
	helpdbname = global.CurrPath + helpdbname
	//fmt.Println("帮助数据库地址:",logindbname)
	tdb, err := gorm.Open(helpdbtye, helpdbname)
	tools.PanicErr(err, "帮助数据库初始化")
	helpdb = tdb
	if !helpdb.HasTable(&models.Help{}) {
		helpdb.CreateTable(&models.Help{})
	}
	//fmt.Println("帮助数据库初始化完成")
	global.WgDb.Done()
}

//发布帮助
func PublishHelp(help *models.Help)(err error){
	u,err:=GetUserById(help.HelperId)
	if err != nil {
		return
	}
	_,err=GetSeekHelp(help.SeekHelpId)
	if err != nil {
		return
	}
	help.Time=time.Now()
	help.HelperName=u.Name
	help.HelperHead=u.Head
	help.Accept=-1
	err=helpdb.Create(help).Error
	return
}

//找到帮助
func GetHelpById(hid int)(help *models.Help,err error){
	help=new(models.Help)
	err=helpdb.First(help,hid).Error
	if err != nil {
		err=global.NoSuchHelp
		return
	}
	return
}

//接受帮助
func AcceptHelp(hid int)(err error){
	help,err:=GetHelpById(hid)
	if err != nil {
		return
	}
	//获取求助消息
	sh,err:=GetSeekHelp(help.SeekHelpId)
	if err != nil {
		return
	}
	if sh.Status==1{
		err=global.AlreadyAccept
		return
	}
	//将求助修改为解决
	err=SolveSeekHelp(sh.Id)
	if err != nil {
		return
	}
	//添加帮助者金币
	err=AddUserGold(help.HelperId,sh.Gold)
	if err != nil {
		return
	}
	help.Accept=1
	err=helpdb.Save(help).Error
	if err != nil {
		return
	}
	//发送消息
	err=SendHelpAcceptMsg(sh,help)
	if err != nil {
		return
	}
	return
}

//获取某求助的帮助数量
func GetHelpCount(sid int)(count int,err error){
	err=helpdb.Model(&models.Help{}).Where("seek_help_id = ?", sid).Count(&count).Error
	if err != nil {
		return
	}
	return
}

//查看某求助的帮助
func GetSeekHelpHelps(sid int, limit, page int)(hs *[]models.Help,err error){
	h:=&models.Help{
		SeekHelpId:sid,
	}
	hs=new([]models.Help)
	err=helpdb.Where(h).Offset((page - 1) * limit).Limit(limit).Find(hs).Error
	if err != nil {
		return
	}
	return
}
