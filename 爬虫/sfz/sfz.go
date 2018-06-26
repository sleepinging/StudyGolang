package main

import (
	"twt/nettools"
	"fmt"
	"time"
)

func main() {
	url:="https://jix.122.gov.cn/user/m/user/queryuser"
	data:="yhdh=3304****1337&xm=www&yhlx=1"
	headers:=map[string]string{}
	headers["Referer"]="https://jix.122.gov.cn/views/getpass1.html"
	t:=time.Now()
	res,err:=nettools.HttpPost(url,data,headers)
	fmt.Println(time.Now().Sub(t),res,err)
}
