package main

import (
	"./model"
	"./tool"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

var (
	freq      = time.Second * 30 //查询间隔
	cfg       model.Config
	exp       = regexp.MustCompile(`(?s)(?i)"amount":"(.+?)"`)
	url       = `https://wbs.icbc.com.cn/wbsaccountserviceweb/accountservice/getDebitCardBalance.do`
	header    = map[string]string{}
	data      = `{"cardNumKey": "1dde3f208575762809225216_cardNum_1",
  		"sessionId": "NWIyOGJlMjExODE2NTUxNzQ4MjA2NTkyIzFkZGUzZjIwODU3NTc2MjgwOTIyNTIxNg\u003d\u003d",
  		"AppId": "DZYH"}`
)

func init() {
	header["Referer"] = `https://wbs.icbc.com.cn/wbsaccountserviceweb/Trabsfer`
	header["Content-Type"] = `application/json;charset=UTF-8`
	path,err:=tool.GetCurrentPath()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	err = loadcfg(path+"config.json")
	if err != nil {
		fmt.Println(err)
		fmt.Println("请将配置文件config.json放入此目录:", path)
		panic(err)
	}
}

func loadcfg(filename string) (err error) {
	err = cfg.Load(filename)
	if err != nil {
		fmt.Println("加载配置文件失败，请确保json格式正确")
		return
	}
	fmt.Println("加载配置文件成功")
	tool.InitSmtpCfg(&cfg.SMTP)
	return
}

//获取json结果
func getres() (res string, err error) {
	return tool.HttpPost(url, data, header)
}

//获取余额
func GetMoney() (m float64, err error) {
	res, err := getres()
	if err != nil {
		//fmt.Println(err)
		return
	}
	rs := exp.FindSubmatch([]byte(res))
	if len(rs) > 0 {
		//fmt.Println(string(rs[1]))
		//转数字
		m, err = strconv.ParseFloat(string(rs[1]), 64)
		if err != nil {
			return
		}
	} else {
		err = fmt.Errorf("未从%s中找到金额", res)
		return
	}
	return
}

func StartWatch() {
	lastmoney,err:= GetMoney()
	minmsged:=false
	if err == nil {
		fmt.Println("服务启动成功,当前银行卡余额为", fmt.Sprintf("%.2f", lastmoney), "元")
	}
	for {
		time.Sleep(freq)
		cm,err := GetMoney()
		if err != nil {
			continue
		}
		d := cm - lastmoney
		if math.Abs(d) > 0.005 { //余额变动
			if d > 0 {
				fmt.Println(time.Now())
				fmt.Println("余额增加")
				go tool.SendEmailToMe("银行卡充值提醒",
					"刚才充值"+fmt.Sprintf("%.2f", d)+"元<br>"+
						"充值时间:"+time.Now().Format("2006/01/02 15:04:05")+"<br>"+
						"当前余额:"+fmt.Sprintf("%.2f", cm)+"元")
			} else {
				fmt.Println(time.Now())
				fmt.Println("余额减少")
				go tool.SendEmailToMe("银行卡消费提醒",
					"刚才消费"+fmt.Sprintf("%.2f", -d)+"元<br>"+
						"消费时间:"+time.Now().Format("2006/01/02 15:04:05")+"<br>"+
						"当前余额:"+fmt.Sprintf("%.2f", cm)+"元")
			}
			minmsged = false
		}
		if cm < float64(cfg.Limit) { //小于阈值
			if !minmsged { //如果没有提醒过
				fmt.Println(time.Now())
				fmt.Println("余额过低")
				go tool.SendEmailToMe("银行卡余额提醒",
					"只剩"+fmt.Sprintf("%.2f", cm)+
						"元了,请尽快充值")
				minmsged = true
			}
		} else {
			minmsged = false
		}
		lastmoney = cm
	}
}

func main() {
	StartWatch()
}
