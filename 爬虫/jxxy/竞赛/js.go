package main

import (
	"fmt"
	"github.com/pysrc/bs"
	"twt/nettools"
	."./dao"
)

func Html2SSIs(html string) (ssi []SStuInfo, err error) {
	soup := bs.Init(html)
	sc:=0
	nstus:=soup.SelByClass("maxtable-row")
	ssi=make([]SStuInfo,len(nstus))
	for _, i := range nstus { //每个人
		pss := make([]string, 11)
		c:=0
		for _, ss := range i.Sel("span", &map[string]string{"class": "ct name"}) {
			pss[c]=ss.Value
			c++
		}
		ssi[sc].CopyFromSS(pss)
		sc++
	}
	return
}

func GetSSIsFromId(id int)(ssi []SStuInfo, err error){
	url := `http://jspt.zjxu.edu.cn/admin/ContestGroupMemberList.aspx?id=`+fmt.Sprintf("%d",id)
	headers := map[string]string{}
	headers["Cookie"] = `UserInfo=MDEwOTA4NzA3MDcwNzgwNzc3MDMwNTBEM0UwNjcxNzQwRDA1MDc3NzcxMDQwNzcwNzIwODc2NzkwRjcxNzgwMjAzN0IwRjdDNzU3NzcyNzQwMzAzMEUwOTc2Mzk3Mg%3d%3d`
	res, err := nettools.HttpGeth(url, headers)
	if err != nil {
		return
	}
	ssi,err=Html2SSIs(res)
	if err != nil {
		return
	}
	return
}

func startp(startid int){
	for {
		fmt.Println("id=",startid)
		ss,err:=GetSSIsFromId(startid)
		if err == nil {
			err=SaveSStuInfos(ss)
			if err != nil {
				fmt.Println(err)
			}
		}else {
			fmt.Println(err)
		}
		startid++
		//time.Sleep(time.Second*1)
	}
}

func main() {
	startp(1)
}
