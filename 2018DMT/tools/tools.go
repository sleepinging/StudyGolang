package tools

import (
	"../models"
	"fmt"
	"encoding/json"
	"net/http"
	"time"
	"os"
	"path"
	"strings"
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

func FmtTime() (ftime string) {
	ftime = time.Now().Format("2006-01-02 15:04:05")
	return
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GenFileName(filename string) (newname string) {
	tz := path.Ext(filename)
	newname = strings.TrimSuffix(filename, tz)
	newname += ` ` + time.Now().String()
	newname, _ = Encrypt(newname, "TWT1234567890TWT")
	newname += tz
	return
}
