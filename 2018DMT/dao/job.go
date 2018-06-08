package dao

import (
	"../global"
	"../models"
	"../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
	"time"
)

var (
	jobdbname   = global.Config.DbInfo.JobDb     //工作数据库
	jobdbtye    = global.Config.DbInfo.JobDbType //数据库类型
	jobdb       *gorm.DB                         //数据库连接
	jobtx       *gorm.DB                         //事务
	jobmodified = false                          //被修改
	wgjobop     *sync.WaitGroup                  //正在操作数据库
	wgjobcommit *sync.WaitGroup                  //正在提交事务
)
//TODO 后期加入回滚
func init() {
	global.WgDb.Add(1)
	go JobDbInit()
}

func JobDbInit() {
	jobdbname = global.CurrPath + jobdbname
	tdb, err := gorm.Open(jobdbtye, jobdbname)
	tools.PanicErr(err, "工作数据库初始化")
	jobdb = tdb
	if !jobdb.HasTable(&models.Job{}) {
		jobdb.CreateTable(&models.Job{})
	}
	jobtx = jobdb.Begin()
	wgjobcommit = new(sync.WaitGroup)
	wgjobop = new(sync.WaitGroup)
	go StartAutoJobCommit()
	//fmt.Println("工作数据库初始化完成")
	global.WgDb.Done()
}

func PublishJob(job *models.Job) (id int, err error) {
	//如果需要提交的话先等待提交
	wgjobcommit.Wait()
	//正在操作
	wgjobop.Add(1)

	job.PublishTime = time.Now()
	u, err := GetUserById(job.PublisherId)
	if err != nil {
		return
	}
	job.PublisherName = u.Name
	err = jobtx.Create(job).Error
	id = job.Id

	jobmodified = true
	//完成本次操作
	wgjobop.Done()
	return
}

func ShowJob(id int) (job *models.Job, err error) {
	job = new(models.Job)
	jobdb.Where(&models.Job{Id: id}).First(job)
	if job.Id == 0 {
		err = global.NoSuchJob
		return
	}
	return
}

func QueryJobCount(job *models.Job) (count int) {
	jobs := new([]models.Job)
	jobdb.Where(job).Find(jobs)
	count = len(*jobs)
	return
}

func QueryJob(job *models.Job, limit, page int) (jobs *[]models.Job) {
	jobs = new([]models.Job)
	jobdb.Where(job).Offset((page - 1) * limit).Limit(limit).Find(jobs)
	return
}

func UpdataJob(id int, newjob *models.Job) (err error) {
	//如果需要提交的话先等待提交
	wgjobcommit.Wait()
	//正在操作
	wgjobop.Add(1)

	job := new(models.Job)
	c := 0
	jobtx.Where(&models.Job{Id: id}).First(job).Count(&c)
	if c == 0 {
		err = global.NoSuchJob
		return
	}
	job.CopyJobFromEId(newjob)
	//job.PublishTime=time.Now()
	jobtx.Save(job)

	jobmodified = true
	//完成本次操作
	wgjobop.Done()
	return
}

func DeleteJob(id int) (err error) {
	//如果需要提交的话先等待提交
	wgjobcommit.Wait()
	//正在操作
	wgjobop.Add(1)

	job := new(models.Job)
	jobtx.Where(&models.Job{Id: id}).First(job)
	if job.Id == 0 {
		err = global.NoSuchJob
		return
	}
	jobtx.Delete(job)

	jobmodified = true
	//完成本次操作
	wgjobop.Done()
	return
}

//根据ID查询工作
func GetJobById(jid int) (job *models.Job, err error) {
	qjob := &models.Job{Id: jid}
	job = new(models.Job)
	jobdb.Where(qjob).Find(job)
	if job.PublisherId == 0 {
		err = global.NoSuchJob
		return
	}
	return
}

func StartAutoJobCommit() {
	for {
		time.Sleep(global.Config.DbTransCommitTime)
		if jobmodified {
			//先表示需要提交事务
			wgjobcommit.Add(1)
			//等待其它进程完成
			wgjobop.Wait()

			jobtx.Commit()
			jobmodified = false

			jobtx = jobdb.Begin()
			wgjobcommit.Done()
		}
	}
}