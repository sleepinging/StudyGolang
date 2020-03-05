package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"reflect"
)

type SStuInfo struct {
	Index    string //序号
	Num      string `gorm:"primary_key"`//学号
	Name     string //姓名
	Sex      string //性别
	Phone    string //联系方式
	Email    string //邮箱
	Ins      string //学院
	Class    string //班级
	JionDate string //加入日期
	Captain  string //是否队长
	Status   string //状态
}

func min(a, b int) (m int) {
	if a < b {
		m = a
	} else {
		m = b
	}
	return
}

//全部复制
func (this *SStuInfo) CopyFromSS(ss []string) {
	ve := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*this)
	n := min(typ.NumField(), len(ss))
	for i := 0; i < n; i++ {
		ve.Field(i).Set(reflect.ValueOf(ss[i]))
	}
}

var db *gorm.DB

func init() {
	var err error=nil
	db, err = gorm.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !db.HasTable(&SStuInfo{}) {
		db.CreateTable(&SStuInfo{})
	}
	//fmt.Println("inited!")
}

func SaveSStuInfos(infos []SStuInfo) (err error) {
	tdb:=db.Begin()
	for _,info:=range infos{
		fmt.Println(info)
		err=tdb.Create(info).Error
		if err != nil {
			tdb.Rollback()
			continue
		}
	}
	tdb.Commit()
	return
}

func SaveSStuInfo(info *SStuInfo) (err error) {
	err=db.Create(info).Error
	if err != nil {
		return
	}
	return
}
