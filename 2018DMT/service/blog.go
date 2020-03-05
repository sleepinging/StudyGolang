package service

import (
	"net/http"
	"../dao"
	"../control/Permission"
	"../models"
	"net/url"
	"fmt"
)

//发布博客
func PublishBlog(w http.ResponseWriter, r *http.Request){
	uid, err := Permission.PublishBlog(w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	r.ParseForm()
	blogs,err:=GetPostString("Blog",w,r)
	fmt.Println(blogs)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	blog,err:=models.LoadBlogFromStr(blogs)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	blog.PublisherId=uid
	err=dao.PublishBlog(blog)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", blog, w)
}

//显示博客
func GetBlog(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id,err:=GetGetInt("Id",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	blog,err:=dao.GetBlogById(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	go dao.AddBlogReadedNum(id)
	models.SendRetJson2(1, "成功", blog, w)
}

//搜索博客
func SearchBlog(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit,err:=GetGetInt("Limit",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page,err:=GetGetInt("Page",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	bs,err:=GetGetString("Blog",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	blog,err:=models.LoadBlogFromStr(bs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	blogs,err:=dao.SearchBlog(blog,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", blogs, w)
}

//搜索博客的数量
func CountSearchBlog(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	bs,err:=GetGetString("Blog",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	blog,err:=models.LoadBlogFromStr(bs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	count,err:=dao.CountSearchBlog(blog)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", count, w)
}

//回复博客
func ReplyBlog(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	brs,err:=GetPostString("BlogReply",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	br,err:=models.LoadBlogReplyFromStr(brs)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.ReplyBlog(br,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	err=dao.ReplyBlog(br)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", br, w)
}

//点赞博客
func ZanBlog(w http.ResponseWriter, r *http.Request){
	uid,err:=Permission.Zan(w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	r.ParseForm()
	bid,err:=GetPostInt("BlogId",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err=dao.ZanBlog(bid,uid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", err.Error(), w)
}

//取消赞
func CancelZanBlog(w http.ResponseWriter, r *http.Request){
	uid,err:=Permission.Zan(w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	r.ParseForm()
	bid,err:=GetPostInt("BlogId",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	err=dao.CancelZanBlog(bid,uid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", err.Error(), w)
}

//是否赞过
func CheckZanBlog(w http.ResponseWriter, r *http.Request){
	//这里不需要判断了
	uid,err:=Permission.Zan(w,r)
	r.ParseForm()
	bid,err:=GetPostInt("BlogId",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	f,err:=dao.IsZanBlog(bid,uid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", f, w)
}

//修改博客
func UpdateBlog(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	id,err:=GetPostInt("Id",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.UpDateBlog(id,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	blogs,err:=GetPostString("Blog",w,r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	blog,err:=models.LoadBlogFromStr(blogs)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	err=dao.UpdateBlog(id,blog)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", blog, w)
}

//删除博客
func DeleteBlog(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	id,err:=GetPostInt("Id",w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=Permission.DeleteBlog(id,w,r)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	_=uid
	err=dao.DeleteBlog(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", "(⊙o⊙)？", w)
}

//获取博客回复的数量
func CountBlogReply(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	bid,err:=GetGetInt("BlogId",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	c,err:=dao.CountBlogReply(bid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
	return
}

//获取博客回复
func GetBlogReply(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	bid,err:=GetGetInt("BlogId",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit,err:=GetGetInt("Limit",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page,err:=GetGetInt("Page",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	brs,err:=dao.GetBlogReply(bid,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", brs, w)
	return
}

//查找某人发布的博客
func GetUserBlogs(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	id,err:=GetGetInt("Id",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	bs,err:=dao.GetUserBlog(id)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "cg", bs, w)
	return
}

//博客发布量
func BlogPublishCount(w http.ResponseWriter, r *http.Request)  {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	d,err:=GetGetInt("Day",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	c,err:=dao.BlogPublishCount(d)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}

//查看某用户赞过的博客
func GetUserZanBlogs(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=GetGetInt("UserId",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	limit,err:=GetGetInt("Limit",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	page,err:=GetGetInt("Page",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	res,err:=dao.GetUserZanBlogs(uid,limit,page)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", res, w)
}

//查看某用户赞过的博客数
func CountUserZanBlogs(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	uid,err:=GetGetInt("UserId",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	res,err:=dao.CountUserZanBlogs(uid)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", res, w)
}