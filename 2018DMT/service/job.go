package service

import (
	"net/http"
	"../dao"
	"../models"
	"../control/Permission"
	"encoding/json"
	"net/url"
	"strconv"
)

func PublishJob(w http.ResponseWriter, r *http.Request) {
	uid, err := Permission.PublishJob(w, r)
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok || len(jobs) == 0 {
		models.SendRetJson2(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson2(0, "Job格式错误", err.Error(), w)
		return
	}
	job.PublisherId = uid
	id, err := dao.PublishJob(job)
	if err != nil {
		models.SendRetJson2(0, "发布失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "发布成功", id, w)
	return
}

func ShowJob(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	if len(queryForm["Id"]) == 0 {
		models.SendRetJson2(0, "缺少Id参数", "手动滑稽", w)
		return
	}
	id, err := strconv.ParseInt(queryForm["Id"][0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "id参数需要为整数", "手动滑稽", w)
		return
	}
	job, err := dao.ShowJob(int(id))
	if err != nil {
		models.SendRetJson2(0, "错误", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", job, w)
	return
}

func QueryJobCount(w http.ResponseWriter, r *http.Request) {
	f, err := Permission.QueryJob(w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	jobs,err:=GetGetString("Job",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs), job)
	if err != nil {
		models.SendRetJson2(0, "Job格式错误", err.Error(), w)
		return
	}
	c := dao.QueryJobCount(job)
	models.SendRetJson2(1, "查询成功", c, w)
	return
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	f, err := Permission.QueryJob(w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
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
	jobs,err:=GetGetString("Job",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs), job)
	if err != nil {
		models.SendRetJson2(0, "Job格式错误", err.Error(), w)
		return
	}
	qjobs := dao.QueryJob(job, int(limit), int(page))
	models.SendRetJson2(1, "成功", qjobs, w)
	return
}

func UpdataJob(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok || len(jobs) == 0 {
		models.SendRetJson2(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson2(0, "缺少Id参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err := json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson2(0, "Job格式错误", err.Error(), w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	err = dao.UpdataJob(int(id), job)
	if err != nil {
		models.SendRetJson2(0, "修改失败", err.Error(), w)
		return
	}
	f, err := Permission.UpdateJob(int(id), w, r)
	if !f {
		models.SendRetJson2(0, "", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "修改成功", ids[0], w)
	return
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson2(0, "缺少id参数", "手动滑稽", w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson2(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	err = dao.DeleteJob(int(id))
	if err != nil {
		models.SendRetJson2(0, "删除失败", err.Error(), w)
		return
	}
	f, err := Permission.DeleteJob(int(id), w, r)
	if !f {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "删除成功", ids[0], w)
	return
}

//工作发布量
func JobPublishCount(w http.ResponseWriter, r *http.Request)  {
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
	c,err:=dao.JobPublishCount(d)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}
