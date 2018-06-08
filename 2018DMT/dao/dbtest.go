package dao

import (
	"../models"
	"fmt"
)

func QueryTest(){
	c:=0
	us:=new([]models.User)
	userdb.
		Where(&models.User{Sex:1}).
		Find(us).
		Where(`name like ?`,`%测试%`).
		Or(`hometown like ?`,`%测试%`).
		Find(us).Count(&c)
	//userdb.Where("sex >= ?",1).Find(us)
	fmt.Println(c,us)
}
