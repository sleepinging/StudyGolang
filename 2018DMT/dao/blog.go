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
	if !blogdb.HasTable(&models.BlogZan{}) {
		blogdb.CreateTable(&models.BlogZan{})
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
	blog.ZanNum=0
	blog.ReadNum=0
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
func SearchBlog(blog *models.Blog, limit, page int)(blogs *[]models.Blog,err error){
	blogs=new([]models.Blog)
	b:=&models.Blog{
		Type:blog.Type,
	}
	err=blogdb.Model(b).Where("title like ? or content like ?",
		`%`+blog.Title+`%`,`%`+blog.Title+`%`).Where(b).
		Offset((page - 1) * limit).Limit(limit).
		Order("time desc").
		Find(blogs).Error
	if err != nil {
		return
	}
	return
}

//搜索博客的数量
func CountSearchBlog(blog *models.Blog)(count int,err error){
	b:=&models.Blog{
		Type:blog.Type,
	}
	err=blogdb.Model(b).Where("title like ? or content like ?",
		`%`+blog.Title+`%`,`%`+blog.Title+`%`).Where(b).
		Count(&count).Error
	if err != nil {
		return
	}
	return
}

//点赞
func ZanBlog(bid,uid int)(err error){
	u,err:=GetUserById(uid)
	if err != nil {
		return
	}
	b,err:=GetBlogById(bid)
	if err != nil {
		return
	}
	zb:=&models.BlogZan{
		BlogId:bid,
		ZanerId:uid,
		ZanerName:u.Name,
	}
	err=blogdb.Create(zb).Error
	if err != nil {
		err=global.AlreadyZan
		return
	}
	b.ZanNum++
	err=blogdb.Save(b).Error
	if err != nil {
		return
	}
	return
}

//取消点赞
func CancelZanBlog(bid,uid int)(err error){
	b,err:=GetBlogById(bid)
	if err != nil {
		return
	}
	zb:=&models.BlogZan{
		BlogId:bid,
		ZanerId:uid,
	}
	tb:=zb
	err=blogdb.Where(zb).Find(tb).Error
	if err != nil {
		return global.HaventZan
	}
	blogdb.Delete(tb)
	b.ZanNum--
	err=blogdb.Save(b).Error
	if err != nil {
		return
	}
	return
}

//是否已赞
func IsZanBlog(bid,uid int)(c int,err error){
	err=blogdb.Model(&models.BlogZan{}).Where("blog_id = ? and zaner_id =?",
		bid,uid).
		Count(&c).Error
	return
}

//添加阅读数
func AddBlogReadedNum(bid int){
	b:=new(models.Blog)
	blogdb.Select("read_num").First(b,bid)
	b.ReadNum++
	blogdb.Model(b).Where("id =?",bid).Updates(b)
}

//查找某人发布的博客
func GetUserBlog(uid int)(bs *[]models.Blog,err error){
	bs=new([]models.Blog)
	err=blogdb.Model(&models.Blog{}).
		Where("publisher_id =?",uid).
		Select([]string{"id", "title"}).
		Find(bs).Error
	if err != nil {
		return
	}
	return
}
