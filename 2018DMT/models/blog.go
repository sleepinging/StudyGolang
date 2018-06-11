package models

import (
	"../tools"
	"reflect"
	"time"
)

//博客
type Blog struct {
	Id            int       `gorm:"primary_key" gorm:"AUTO_INCREMENT" json:"Id"`
	PublisherId   int       `json:"PublisherId"`
	PublisherName string    `json:"PublisherName"` //+
	PublisherHead string    `json:"PublisherHead"` //+
	Time          time.Time `json:"Time"`
	Status        int       `json:"Status"`
	Type          int       `json:"Type"`
	Title         string    `json:"Title"`
	Content       string    `gorm:"type:varchar(10239)" json:"Content"`
	ReadNum       int       `json:"ReadNum"`
	ZanNum        int       `json:"ZanNum"`   //+
	ReplyNum      int       `json:"ReplyNum"` //+
}

//除了某些字段全部复制
func (this *Blog) CopyBlogFromExpt(blog *Blog, except []string) {
	s1 := reflect.ValueOf(blog).Elem()
	s2 := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*blog)
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
