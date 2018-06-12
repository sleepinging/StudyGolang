package service

import (
	"net/http"
	"../dao"
	"../control/Permission"
	"../models"
	"net/url"
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

}

//点赞博客
func ZanBlog(w http.ResponseWriter, r *http.Request){

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
	models.SendRetJson2(1, "成功", err.Error(), w)
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
