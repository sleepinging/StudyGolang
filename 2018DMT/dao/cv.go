package dao

import (
	"../global"
	"../models"
	"../tools"
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

var (
	cvdbname   = global.Config.DbInfo.CVDb     //用户数据库
	cvdbtye    = global.Config.DbInfo.CVDbType //数据库类型
	cvdb       *gorm.DB                        //数据库连接
	cvtx       *gorm.DB                        //事务
	cvmodified = false                         //被修改
	wgcvop     *sync.WaitGroup                 //正在操作数据库
	wgcvcommit *sync.WaitGroup                 //正在提交事务
)

func init() {
	global.WgDb.Add(1)
	go CVDbInit()
}

//初始化
func CVDbInit() {
	cvdbname = global.CurrPath + cvdbname
	//fmt.Println("用户数据库地址:",logindbname)
	tdb, err := gorm.Open(cvdbtye, cvdbname)
	tools.PanicErr(err, "简历数据库初始化")
	cvdb = tdb
	if !cvdb.HasTable(&models.CV{}) {
		cvdb.CreateTable(&models.CV{})
	}
	cvtx = cvdb.Begin()
	wgcvcommit = new(sync.WaitGroup)
	wgcvop = new(sync.WaitGroup)
	go StartAutoCVCommit()
	fmt.Println("简历数据库初始化完成")
	global.WgDb.Done()
}

//获取用户简历
func GetUserCV(uid int) (cvs *[]models.CV, err error) {
	cvs = new([]models.CV)
	qcv := &models.CV{UserId: uid}
	err = cvdb.Where(qcv).Find(cvs).Error
	if len(*cvs) == 0 {
		err = global.UserNoCV
		return
	}
	return
}

//获取简历
func GetCV(cid int) (cv *models.CV, err error) {
	cv = new(models.CV)
	qcv := &models.CV{Id: cid}
	err = cvdb.Where(qcv).Find(cv).Error
	if cv.Id == 0 {
		err = global.NoSuchCV
		return
	}
	return
}

//修改简历
func UpdateCV(cid int, cv *models.CV) (err error) {
	//如果需要提交的话先等待提交
	wgcvcommit.Wait()
	//正在操作
	wgcvop.Add(1)

	qcv := &models.CV{Id: cid}
	ocv := new(models.CV)
	cvtx.Where(qcv).First(ocv)
	if ocv.Id == 0 {
		err = global.NoSuchCV
		return
	}
	ocv.UserId = cv.UserId
	ocv.Context = cv.Context
	err = cvtx.Save(ocv).Error

	cvmodified = true
	//完成本次操作
	wgcvop.Done()
	return
}

//删除简历
func DeleteCV(cid int) (err error) {
	//如果需要提交的话先等待提交
	wgcvcommit.Wait()
	//正在操作
	wgcvop.Add(1)

	qcv := &models.CV{Id: cid}
	cv := new(models.CV)
	cvtx.Where(qcv).First(cv)
	if cv.Id == 0 {
		err = global.NoSuchCV
		return
	}
	err = cvtx.Delete(cv).Error

	cvmodified = true
	//完成本次操作
	wgcvop.Done()
	return
}

//添加简历
func AddCV(cv *models.CV) (cid int, err error) {
	//如果需要提交的话先等待提交
	wgcvcommit.Wait()
	//正在操作
	wgcvop.Add(1)

	err = cvtx.Create(cv).Error
	cid = cv.Id

	cvmodified = true
	//完成本次操作
	wgcvop.Done()
	return
}

func StartAutoCVCommit() {
	for {
		time.Sleep(global.Config.DbTransCommitTime)
		if cvmodified {
			//先表示需要提交事务
			wgcvcommit.Add(1)
			//等待其它进程完成
			wgcvop.Wait()

			cvtx.Commit()
			cvmodified = false

			cvtx = cvdb.Begin()
			wgcvcommit.Done()
		}
	}
}
