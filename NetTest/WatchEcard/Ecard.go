package main

import (
	"./model"
	"./tool"
	"fmt"
	"math"
	"strconv"
	"time"
)

var (
	lastmoney  = 0.0
	freq            = time.Second * 30 //查询间隔
	cfg       model.Config
	minmsged  = false
)

func GetMoney() (money float64) {
	return 15
	//fmt.Println("查询中")
	res, err := tool.HttpGet(
		"http://wxschool.lsmart.cn/card/queryAcc_queryAccount.shtml?" +
			"openId=" + cfg.Openid +
			"&wxArea=10354")
	if err != nil {
		panic(err)
	}
	smoney, err := tool.GetBetween(res, `校园卡余额<strong>`, `</strong>元`)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	money, _ = strconv.ParseFloat(smoney, 32) //转化为数字
	return
}
func StartWatch() {
	lastmoney := GetMoney()
	fmt.Println("服务启动成功,当前校园卡余额为", fmt.Sprintf("%.2f", lastmoney), "元")
	for {
		time.Sleep(freq)
		cm := GetMoney()
		if cm <= float64(cfg.Limit) && !minmsged {
			go tool.SendEmailToMe("校园卡余额提醒",
				"只剩" + fmt.Sprintf("%.2f", cm)+
					"元了,请尽快充值")
			minmsged = true //已经提醒过了
		} else {
			minmsged = false
		}
		d := cm - lastmoney
		if math.Abs(d) > 0.01 {
			if d > 0 {
				fmt.Println("余额增加")
				go tool.SendEmailToMe("校园卡充值提醒",
					"刚才充值" + fmt.Sprintf("%.2f", d) + "元<br>"+
						"当前余额"+ fmt.Sprintf("%.2f", cm)+ "元")
			} else {
				fmt.Println("余额减少")
				go tool.SendEmailToMe("校园卡消费提醒",
					"刚才消费" + fmt.Sprintf("%.2f", d) + "元<br>"+
						"当前余额"+ fmt.Sprintf("%.2f", cm)+ "元")
			}
		}
		lastmoney = cm
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
func main() {
	fmt.Println("本项目免费开源",
		"地址：https://github.com/sleepinging/StudyGolang/tree/master/NetTest/%E8%8E%B7%E5%8F%96%E6%A0%A1%E5%9B%AD%E5%8D%A1%E4%BD%99%E9%A2%9D")
	err := loadcfg(`config.json`)
	if err != nil {
		fmt.Println(err)
		path, _ := tool.GetCurrentPath()
		fmt.Println("请将配置文件config.json放入此目录:", path)
		return
	}
	StartWatch()
	//fmt.Println(cfg)
	//fmt.Println(GetMoney())
}
