package main

import (
	"twt/nettools"
	"fmt"
)

func logineh(){
	url:=`http://www.effecthub.com/login/check`
	hs:=map[string]string{}
	hs["Referer"]=`http://www.effecthub.com/login`
	hs["Origin"]=`http://www.effecthub.com`
	hs["Accept-Language"]=`zh-CN,zh;q=0.9`
	hs["Accept"]=`text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8`
	hs["User-Agent"]=`Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`
	data:=`email=237731947%40qq.com&password=eh123456&remember_me=on&redirectURL=`
	fmt.Println(nettools.HttpPost(url,data,hs))
}

func main() {
	logineh()
}
