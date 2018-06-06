package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Job struct {
	Id          int     `gorm:"primary_key"`
	PublisherId int     `json:"PublisherId"`
	Name        string  `json:"Name"`
	Salary      float32 `json:"Salary"`
	Time        string  `json:"Time"`
	Weekend     int     `json:"Weekend"`
	Pickup      int     `json:"Pickup"`
	Eat         int     `json:"Eat"`
	Live        int     `json:"Live"`
	WuXianYiJin int     `json:"WuXianYiJin"`
	Place       string  `json:"Place"`
	LimPeople   int     `json:"LimPeople"`
	NowPeople   int     `json:"NowPeople"`
	Sex         int     `json:"Sex"`
	Phone       string  `json:"Phone"`
	Detail      string  `gorm:"type:varchar(2047)" json:"Detail"`
}

//除了id全部复制
func (this *Job) CopyJobFromEId(job *Job) {
	//this.Name = job.Name
	//this.Salary = job.Salary
	//this.Time = job.Time
	//this.Weekend = job.Weekend
	//this.Pickup = job.Pickup
	//this.Eat = job.Eat
	//this.Live = job.Live
	//this.WuXianYiJin = job.WuXianYiJin
	//this.Place = job.Place
	//this.LimPeople = job.LimPeople
	//this.NowPeople = job.NowPeople
	//this.Sex = job.Sex
	//this.Phone = job.Phone
	//this.Detail = job.Detail
	s1 := reflect.ValueOf(job).Elem()
	s2 := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*job)
	n := typ.NumField()
	for i := 0; i < n; i++ {
		name := s1.Type().Field(i).Name
		if name == "Id" {
			continue
		}
		v := s1.Field(i).Interface()
		s2.Field(i).Set(reflect.ValueOf(v))
		//switch v.(type) {
		//case string:
		//	s2.Field(i).SetString(v.(string))
		//case int:
		//	s2.Field(i).SetInt(int64(v.(int)))
		//case float32:
		//	s2.Field(i).SetFloat(float64(v.(float32)))
		//}
	}
}

type JobTypeDess struct {
	Types []JobTypeDes
}

type JobTypeDes struct {
	Description string
	Types       []JobType2
}

type JobType2 struct {
	Name   string
	TypeId string
	Types  []JobType
}

type JobType struct {
	JobId   int
	JobName string
}

func (this *JobTypeDess) Load(filename string) (err error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeType)
	defer file.Close()
	if err != nil {
		err = fmt.Errorf("打开配置文件失败:" + err.Error())
		return
	}
	bs, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("读取配置文件失败:" + err.Error())
		return
	}
	err = json.Unmarshal(bs, this)
	if err != nil {
		err = fmt.Errorf("解析配置文件失败:" + err.Error())
		return
	}
	return
}
