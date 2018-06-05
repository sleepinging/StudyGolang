package models

import (
	"net/http"
	"encoding/json"
	"fmt"
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

//帮助
type Help struct {
	Id int `gorm:"primary_key" json:"Id"`
}

//博客
type Blog struct {
	Id int `gorm:"primary_key" json:"Id"`
}
