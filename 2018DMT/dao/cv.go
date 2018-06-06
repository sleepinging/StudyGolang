package dao

import (
	"github.com/jinzhu/gorm"
	"../global"
	"../models"
	"../tools"
	"fmt"
)

var (
	cvdbname = global.Config.DbInfo.CVDb     //用户数据库
	cvdbtye  = global.Config.DbInfo.CVDbType //数据库类型
	cvdb     *gorm.DB                        //数据库连接
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
	qcv := &models.CV{Id: cid}
	ocv := new(models.CV)
	cvdb.Where(qcv).First(ocv)
	if ocv.Id == 0 {
		err = global.NoSuchCV
		return
	}
	ocv.UserId = cv.UserId
	ocv.Context = cv.Context
	err = cvdb.Save(ocv).Error
	return
}

//删除简历
func DeleteCV(cid int) (err error) {
	qcv := &models.CV{Id: cid}
	cv := new(models.CV)
	cvdb.Where(qcv).First(cv)
	if cv.Id == 0 {
		err = global.NoSuchCV
		return
	}
	err = cvdb.Delete(cv).Error
	return
}

//添加简历
func AddCV(cv *models.CV) (cid int, err error) {
	err = cvdb.Create(cv).Error
	if err != nil {
		return
	}
	sql := `SELECT last_insert_rowid() as id;`
	var lid tid
	err = cvdb.Raw(sql).Select("id").Scan(&lid).Error
	if err != nil {
		return
	}
	cid = lid.Id
	return
}
