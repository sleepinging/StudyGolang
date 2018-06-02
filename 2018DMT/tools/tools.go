package tools

import (
	"../models"
	"fmt"
	"encoding/json"
	"net/http"
)

func SendRetJson(status int, msg, data string, w http.ResponseWriter) {
	retjs := models.RetJson{Status: status, Msg: msg, Data: data}
	res, err := json.Marshal(&retjs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(res))
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func PanicErr(err error, msg string) {
	if err != nil {
		panic(err)
		fmt.Println(msg, "错误")
	}
}
