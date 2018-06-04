package models

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type RetJson struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

type Login struct {
	Email    string `gorm:"primary_key" json:"email"`
	Password string `json:"password"`
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
