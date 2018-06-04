package dao

import (
	"../global"
	"../models"
	"../tools"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"twt/mytools"
)

var (
	jobdbname = global.Config.DbInfo.JobDb     //登录数据库
	jobdbtye  = global.Config.DbInfo.JobDbType //数据库类型
	jobdb     *gorm.DB                         //数据库连接
)

func init() {
	pt, _ := mytools.GetCurrentPath()
	jobdbname = pt + jobdbname
	tdb, err := gorm.Open(jobdbtye, jobdbname)
	tools.CheckErr(err)
	jobdb = tdb
	if !jobdb.HasTable(&models.Job{}) {
		jobdb.CreateTable(&models.Job{})
	}
	fmt.Println("工作数据库初始化完成")
}

type tid struct {
	Id int
}

func PublishJob(job *models.Job) (id int, err error) {
	jobdb.Create(job)
	sql := `SELECT last_insert_rowid() as id;`
	var lid tid
	jobdb.Raw(sql).Select("id").Scan(&lid)
	id = lid.Id
	return
}

func ShowJob(id int) (job *models.Job, err error) {
	job = new(models.Job)
	jobdb.Where(&models.Job{Id: id}).First(job)
	if job.Name == "" {
		err = global.NoSuchJob
		return
	}
	return
}

func QueryJobCount(job *models.Job) (count int) {
	jobs := new([]models.Job)
	jobdb.Where(job).Find(jobs).Count(&count)
	return
}

func QueryJob(job *models.Job, limit, page int) (jobs *[]models.Job) {
	jobs = new([]models.Job)
	jobdb.Where(job).Offset((page - 1) * limit).Limit(limit).Find(jobs)
	return
}
