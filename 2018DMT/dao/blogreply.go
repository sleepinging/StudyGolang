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
	brdbname = global.Config.DbInfo.BlogReplyDb     //用户数据库
	brdbtye  = global.Config.DbInfo.BlogReplyDbType //数据库类型
	brdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go BlogReplyDbInit()
}

//初始化
func BlogReplyDbInit() {
	brdbname = global.CurrPath + brdbname
	//fmt.Println("博客数据库地址:",logindbname)
	tdb, err := gorm.Open(brdbtye, brdbname)
	tools.PanicErr(err, "博客数据库初始化")
	brdb = tdb
	if !brdb.HasTable(&models.BlogReply{}) {
		brdb.CreateTable(&models.BlogReply{})
	}
	//fmt.Println("博客数据库初始化完成")
	global.WgDb.Done()
}

//回复博客
func ReplyBlog(br *models.BlogReply)(err error){
	b,err:=GetBlogById(br.BlogId)
	if err != nil {
		return
	}
	ru,err:=GetUserById(br.ReplyerId)
	if err != nil {
		return
	}
	br.ReplyerName=ru.Name
	br.ReplyerHead=ru.Head
	au,err:=GetUserById(br.AtId)
	if err != nil {
		return
	}
	br.AtName=au.Name
	br.Time=time.Now()
	err=brdb.Create(br).Error
	if err != nil {
		return
	}
	b.ReplyNum++
	err=blogdb.Save(b).Error
	if err != nil {
		return
	}
	return
}

//获取博客回复数量
func CountBlogReply(bid int)(count int,err error){
	err=brdb.Model(&models.BlogReply{}).
		Where("blog_id =?",bid).
		Count(&count).Error
	return
}

//获取博客的回复
func GetBlogReply(bid int, limit, page int)(brs *[]models.BlogReply,err error){
	brs=new([]models.BlogReply)
	err=brdb.Model(&models.BlogReply{}).Where("blog_id =?",
		bid).
		Offset((page - 1) * limit).Limit(limit).
		Order("time desc").
		Find(brs).Error
	return
}
