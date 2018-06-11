package dao


import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"../global"
	"../models"
	"../tools"
)

var (
	blogdbname = global.Config.DbInfo.BlogDb     //用户数据库
	blogdbtye  = global.Config.DbInfo.BlogDbType //数据库类型
	blogdb     *gorm.DB                          //数据库连接
)

func init() {
	global.WgDb.Add(1)
	go BlogDbInit()
}

//初始化
func BlogDbInit() {
	blogdbname = global.CurrPath + blogdbname
	//fmt.Println("博客数据库地址:",logindbname)
	tdb, err := gorm.Open(blogdbtye, blogdbname)
	tools.PanicErr(err, "博客数据库初始化")
	blogdb = tdb
	if !blogdb.HasTable(&models.Blog{}) {
		blogdb.CreateTable(&models.Blog{})
	}
	//fmt.Println("博客数据库初始化完成")
	global.WgDb.Done()
}

//发布博客

//删除博客

//修改博客

//查找博客

//搜索博客
