package main

import (
	"github.com/lxn/win"
	"twt/mystr"
	"time"
	"fmt"
)

func findwindow(wname string) (f bool) {
	r := win.FindWindow(nil, mystr.Str2ui16p(wname))
	return r > 0
}
func main() {
	for {
		time.Sleep(time.Second*1)
		if !findwindow("go_build_main_go.exe"){
			fmt.Println("进程挂了，重启一下，滑稽")
		}
	}
}
