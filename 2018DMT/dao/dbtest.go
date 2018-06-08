package dao

import (
	"../models"
	"fmt"
)

func QueryTest(){
	us:=new([]models.User)
	userdb.
		Where(&models.User{Sex:1}).
		Find(us).
		Where(`name like ?`,`%测试%`).
		Or(`hometown like ?`,`%测试%`).
		Find(us)
	fmt.Println(us)
}
