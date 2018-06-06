package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//返回的json
type RetJson struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func SendRetJson(status int, msg, data string, w http.ResponseWriter) {
	retjs := RetJson{Status: status, Msg: msg, Data: data}
	res, err := json.Marshal(&retjs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(res))
}

//登录的模型
type Login struct {
	Email    string `gorm:"primary_key" json:"email"`
	Password string `json:"password"`
}

//博客点赞表
type BlogZan struct {
	BlogId int `gorm:"primary_key" json:"BlogId"`
	UserId int `gorm:"primary_key" json:"UserId"`
}

//简历
type CV struct {
	Id      int    `gorm:"primary_key"`
	UserId  int    `json:"UserId"`
	Context string `gorm:"type:varchar(2047)" json:"Context"`
}
