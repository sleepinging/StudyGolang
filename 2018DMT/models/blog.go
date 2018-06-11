package models

import (
	"../tools"
	"reflect"
	"time"
	"encoding/json"
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

//创建博客
func NewBlog(pid,tp int,title,content string,status int)(blog *Blog){
	blog=&Blog{
		PublisherId:pid,
		Type:tp,
		Title:title,
		Content:content,
		Status:status,
	}
	return
}

//除了某些字段全部复制
func (this *Blog) CopyFromExpt(blog *Blog, except []string) {
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

//只复制某些字段
func (this *Blog) CopyFrom(blog *Blog, except []string) {
	s1 := reflect.ValueOf(blog).Elem()
	s2 := reflect.ValueOf(this).Elem()
	typ := reflect.TypeOf(*blog)
	n := typ.NumField()
	for i := 0; i < n; i++ {
		name := s1.Type().Field(i).Name
		if f, _ := tools.StrInArray(except, name); f {
			v := s1.Field(i).Interface()
			s2.Field(i).Set(reflect.ValueOf(v))
		}
	}
}

//从字符串加载博客
func LoadBlogFromStr(str string) (blog *Blog, err error) {
	blog = new(Blog)
	err = json.Unmarshal([]byte(str), blog)
	return
}