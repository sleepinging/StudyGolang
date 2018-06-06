package models

import (
	"../tools"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

type User struct {
	Id         int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	Email      string    `json:"Email"`
	Name       string    `json:"Name"`
	Head       string    `json:"Head"`
	Sex        int       `json:"Sex"`
	Birthday   time.Time `json:"Birthday"`
	Hometown   string    `json:"Hometown"`
	Hobby      string    `json:"Hobby"`
	Speciality string    `json:"Speciality"`
	Phone      string    `json:"Phone"`
	QQ         string    `json:"QQ"`
	WeiXin     string    `json:"WeiXin"`
	School     string    `json:"School"`
	Major      string    `json:"Major"`
	Type       int       `json:"Type"`
}

func (this *User) ToString() (str string) {
	b, err := json.Marshal(this)
	tools.ShowErr(err)
	str = string(b)
	return
}

func (this *User) IdToString() (sid string) {
	sid = fmt.Sprintf("%d", this.Id)
	return
}

//除了某些字段全部复制
func (this *User) CopyUserFromExpt(user *User, except []string) {
	s1 := reflect.ValueOf(user).Elem()
	s2 := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*user)
	n := typ.NumField()
	for i := 0; i < n; i++ {
		name := s1.Type().Field(i).Name
		if f, _ := tools.StrInArray(except, name); f {
			continue
		}
		v := s1.Field(i).Interface()
		s2.Field(i).Set(reflect.ValueOf(v))
	}
}
