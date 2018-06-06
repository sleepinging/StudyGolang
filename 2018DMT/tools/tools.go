package tools

import (
	"fmt"
	"time"
	"os"
	"path"
	"strings"
	"path/filepath"
)

func ShowErr(err error) {
	if err != nil {
		//31表示红色
		fmt.Printf("%c[1;0;31m%s%c[0m\n", 0x1b, err.Error(), 0x1b)
		//fmt.Println(err)
	}
}

func PanicErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, "错误")
		panic(err)
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

func GetCurrentPath() (path string, err error) {
	fullname, err := filepath.Abs(os.Args[0])
	PanicErr(err, "获取程序路径")
	fn := filepath.Base(fullname)
	path = strings.TrimSuffix(fullname, fn)
	return
}

func StrInArray(arr []string, e string) (f bool, index int) {
	v := e
	for index, v = range arr {
		if v == e {
			f = true
			return
		}
	}
	return
}