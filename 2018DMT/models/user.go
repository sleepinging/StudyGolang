package models

import (
	"../tools"
	"encoding/json"
	"time"
)

type User struct {
	Id         int       `gorm:"primary_key" json:"Id"`
	Email      string    `json:"Email"`
	Name       string    `json:"Name"`
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
	Gold       int       `json:"Gold"`
	Type       int       `json:"Type"`
}

func (this *User) ToString() (str string) {
	b, err := json.Marshal(this)
	tools.ShowErr(err)
	str = string(b)
	return
}
