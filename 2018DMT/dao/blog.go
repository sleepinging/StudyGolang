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
func PublishBlog(blog *models.Blog)(err error){
	u,err:=GetUserById(blog.PublisherId)
	if err != nil {
		return
	}
	blog.PublisherName=u.Name
	blog.PublisherHead=u.Head
	blog.Time=time.Now()
	err=blogdb.Create(blog).Error
	if err != nil {
		return
	}
	return
}

//删除博客
func DeleteBlog(bid int)(err error){
	blog:=new(models.Blog)
	err=blogdb.First(blog,bid).Error
	if err != nil {
		err=global.NoSuchBlog
		return
	}
	err=blogdb.Delete(blog).Error
	if err != nil {
		return
	}
	return
}

//修改博客
func UpdateBlog(bid int,blog *models.Blog)(err error){
	oblog:=new(models.Blog)
	err=blogdb.First(oblog,bid).Error
	if err != nil {
		return
	}
	oblog.CopyFrom(blog,[]string{"Content","Title","Type","Status"})
	oblog.Time=time.Now()
	err=blogdb.Save(oblog).Error
	if err != nil {
		return
	}
	return
}

//查找博客
func GetBlogById(bid int)(blog *models.Blog,err error){
	blog=new(models.Blog)
	err=blogdb.First(blog,bid).Error
	if err != nil {
		err=global.NoSuchBlog
		return
	}
	return
}

//搜索博客

//搜索博客的数量

//点赞

//取消点赞

//是否已赞
