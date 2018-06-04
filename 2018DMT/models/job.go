package models

import (
	"reflect"
)

//除了id全部复制
func (this *Job) CopyFormEId(job *Job) {
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
		switch v.(type) {
		case string:
			s2.Field(i).SetString(v.(string))
		case int:
			s2.Field(i).SetInt(int64(v.(int)))
		case float32:
			s2.Field(i).SetFloat(float64(v.(float32)))
		}
	}
}
