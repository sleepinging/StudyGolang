package service

import (
	"net/http"
	"../dao"
	"../models"
	"../control"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func PublishJob(w http.ResponseWriter, r *http.Request) {
	f, err := control.PublishJobPermission(w, r)
	if !f {
		models.SendRetJson(0, "发布失败", err.Error(), w)
		return
	}
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok {
		models.SendRetJson(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson(0, "Job格式错误", err.Error(), w)
		return
	}
	id, err := dao.PublishJob(job)
	if err != nil {
		models.SendRetJson(0, "发布失败", err.Error(), w)
		return
	}
	models.SendRetJson(1, "发布成功", fmt.Sprintf("%d", id), w)
	return
}

func ShowJob(w http.ResponseWriter, r *http.Request) {
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson(0, "错误", err.Error(), w)
		return
	}
	if len(queryForm["id"]) == 0 {
		models.SendRetJson(0, "缺少id参数", "手动滑稽", w)
		return
	}
	id, err := strconv.ParseInt(queryForm["id"][0], 10, 32)
	if err != nil {
		models.SendRetJson(0, "id参数需要为整数", "手动滑稽", w)
		return
	}
	job, err := dao.ShowJob(int(id))
	jb, err := json.Marshal(job)
	if err != nil {
		models.SendRetJson(0, "服务器错误", err.Error(), w)
		return
	}
	models.SendRetJson(1, "成功", string(jb), w)
	return
}

func QueryJobCount(w http.ResponseWriter, r *http.Request) {
	f, err := control.QueryJobPermission(w, r)
	if !f {
		models.SendRetJson(0, "查询失败", err.Error(), w)
		return
	}
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok {
		models.SendRetJson(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson(0, "Job格式错误", err.Error(), w)
		return
	}
	c := dao.QueryJobCount(job)
	models.SendRetJson(1, "查询成功", fmt.Sprintf("%d", c), w)
	return
}

func QueryJob(w http.ResponseWriter, r *http.Request) {
	f, err := control.QueryJobPermission(w, r)
	if !f {
		models.SendRetJson(0, "查询失败", err.Error(), w)
		return
	}
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok {
		models.SendRetJson(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err = json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson(0, "Job格式错误", err.Error(), w)
		return
	}
	limits, ok := r.PostForm["Limit"]
	if !ok || len(limits) == 0 {
		models.SendRetJson(0, "缺少limit参数", "手动滑稽", w)
		return
	}
	Pages, ok := r.PostForm["Page"]
	if !ok || len(Pages) == 0 {
		models.SendRetJson(0, "缺少Page参数", "手动滑稽", w)
		return
	}
	limit, err := strconv.ParseInt(limits[0], 10, 32)
	if err != nil {
		models.SendRetJson(0, "Limit参数需要为整数", err.Error(), w)
		return
	}
	page, err := strconv.ParseInt(Pages[0], 10, 32)
	if err != nil {
		models.SendRetJson(0, "Page参数需要为整数", err.Error(), w)
		return
	}
	qjobs := dao.QueryJob(job, int(limit), int(page))
	res, err := json.Marshal(qjobs)
	if err != nil {
		models.SendRetJson(0, "服务器错误", err.Error(), w)
		return
	}
	models.SendRetJson(1, "成功", string(res), w)
	return
}

func UpdataJob(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jobs, ok := r.PostForm["Job"]
	if !ok || len(jobs) == 0 {
		models.SendRetJson(0, "缺少Job参数", "手动滑稽", w)
		return
	}
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson(0, "缺少id参数", "手动滑稽", w)
		return
	}
	job := &models.Job{}
	err := json.Unmarshal([]byte(jobs[0]), job)
	if err != nil {
		models.SendRetJson(0, "Job格式错误", err.Error(), w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	err = dao.UpdataJob(int(id), job)
	if err != nil {
		models.SendRetJson(0, "修改失败", err.Error(), w)
		return
	}
	f, err := control.UpdateJobPermission(int(id), w, r)
	if !f {
		models.SendRetJson(0, "修改失败", err.Error(), w)
		return
	}
	models.SendRetJson(1, "修改成功", ids[0], w)
	return
}

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	ids, ok := r.PostForm["Id"]
	if !ok || len(ids) == 0 {
		models.SendRetJson(0, "缺少id参数", "手动滑稽", w)
		return
	}
	id, err := strconv.ParseInt(ids[0], 10, 32)
	if err != nil {
		models.SendRetJson(0, "Id参数需要为整数", err.Error(), w)
		return
	}
	err = dao.DeleteJob(int(id))
	if err != nil {
		models.SendRetJson(0, "删除失败", err.Error(), w)
		return
	}
	f, err := control.DeleteJobPermission(int(id), w, r)
	if !f {
		models.SendRetJson(0, "删除失败", err.Error(), w)
		return
	}
	models.SendRetJson(1, "删除成功", ids[0], w)
	return
}
