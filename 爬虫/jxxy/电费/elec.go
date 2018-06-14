package main

import (
	"./model"
	"./tool"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	freq = time.Second * 30 //查询间隔
	cfg  model.Config
	msgt = 0 //提醒次数
)

func GetElc() (e float64, err error) {
	data := "openId=" + cfg.Openid +
		"&wxArea=10354&areaNo=1&buildNo=112&levelNo=1001&roomNo=359" +
		"&roomName=%E6%A2%81%E6%9E%97%E5%85%AC%E5%AF%933%E3%80%814%E3%80%815%E3%80%816%E6%A2%81%E6%9E%97%E5%85%AC%E5%AF%935%E5%8F%B7A1%E6%A5%BCLG5-A110"
	res, err := tool.HttpPost(
		"http://wxschool.lsmart.cn/electric/electric_goAmount.shtml",
		data,
		nil)
	if err != nil {
		return
	}
	res, err = tool.GetBetween(res, "<h2>剩余电量</h2>", "</p>")
	if err != nil {
		return
	}
	res, err = tool.GetBetween(res, "<p>", "度")
	if err != nil {
		return
	}
	res = strings.Trim(res, " ")
	e, err = strconv.ParseFloat(res, 32)
	return
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

func Watch() {
	for {
		e, err := GetElc()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if e < cfg.Limit {
			if msgt<=10{
				go Notify(e)
				msgt++
			}
		}else {
			msgt=0
		}
		time.Sleep(freq)
	}
}

func Notify(e float64) {
	fmt.Println("过低", e)
	sub:="寝室电量提醒"
	tool.SendEmailToMe(sub,
		fmt.Sprintf("剩余%.2f度", e))
	time.Sleep(time.Second * 10)
	tool.SendEmailTo(sub,
		fmt.Sprintf("剩余%.2f度", e),
		"1105533959@qq.com")
	time.Sleep(time.Second * 10)
	tool.SendEmailTo(sub,
		fmt.Sprintf("剩余%.2f度", e),
		"1185694080@qq.com")
	time.Sleep(time.Second * 10)
	tool.SendEmailTo(sub,
		fmt.Sprintf("剩余%.2f度", e),
		"1184895114@qq.com")
	time.Sleep(time.Second * 10)
	tool.SendEmailTo(sub,
		fmt.Sprintf("剩余%.2f度", e),
		"3416290468@qq.com")
}

func main() {
	err := loadcfg(`config.json`)
	if err != nil {
		fmt.Println(err)
		path, _ := tool.GetCurrentPath()
		fmt.Println("请将配置文件config.json放入此目录:", path)
		return
	}
	Watch()
	fmt.Println(GetElc())
}
