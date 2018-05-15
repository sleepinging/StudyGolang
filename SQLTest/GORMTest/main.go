package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
)

type User struct {
	//gorm.Model
	Id       int
	Name     string
	Password string
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@/2018dmt?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		return
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	user1 := User{Name: "阿良", Password: "123456"}
	//插入
	db.Create(&user1)

	//删除
	db.Delete(&User{Id: 5})

	// 获取所有匹配记录
	var users []User
	db.Where("name = ?", "阿良").Find(&users)
	// 获取第一个匹配记录
	//db.Where("name = ?", "阿良").First(&users)
	for _, u := range users {
		fmt.Println(u)
	}
}
