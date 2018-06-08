package models

import (
	"../tools"
	"reflect"
	"time"
	"encoding/json"
)

//帮助
type SeekHelp struct {
	Id            int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	PublisherId   int       `json:"PublisherId"`
	PublisherName string    `json:"PublisherName"` //+
	PublisherHead string    `json:"PublisherHead"` //+
	Type          int       `json:"Type"`
	Title         string    `json:"Title"`
	Context       string    `gorm:"type:varchar(4095)" json:"Context"`
	Gold          int       `json:"Gold"`
	Time          time.Time `json:"Time"`
	Status        int       `json:"Status"`
}

//除了某些字段全部复制
func (this *SeekHelp) CopyHelpFromExpt(help *SeekHelp, except []string) {
	s1 := reflect.ValueOf(help).Elem()
	s2 := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*help)
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

//从字符串加载
func LoadSeekHelp(str string)(sh *SeekHelp,err error){
	sh=new(SeekHelp)
	err=json.Unmarshal([]byte(str),sh)
	return
}
