package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
	"../models"
	"../tools"
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

	b.ReplyNum++
	return
}

//获取博客回复数量
func CountBlogReply(bid int)(count int,err error){

	return
}

//获取博客的回复
func GetBlogReply(bid int)(brs *models.BlogReply,err error){
	return
}
