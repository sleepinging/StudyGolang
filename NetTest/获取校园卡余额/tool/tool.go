package tool

import (
	"../model"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"os"
	"os/exec"
	"strings"
)

var (
	smtpcfg *model.Smtp
)

func sendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailtype == "html" {
		contentType = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendto := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, sendto, msg)
	return err
}

func InitSmtpCfg(cfg *model.Smtp) {
	smtpcfg = cfg
}

func SendEmailToMe(subject, text string) (err error) {
	user := smtpcfg.Sender
	password := smtpcfg.Pwd
	host := smtpcfg.Host
	to := smtpcfg.To

	body := `
		<html>
		<body>
		<h3>
		` + text + `
		</h3>
		</body>
		</html>
		`
	fmt.Println("发送邮件...")
	err = sendToMail(user, password, host, to, subject, body, "html")
	return
}
func HttpGet(url string) (retstr string, err error) {
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, strings.NewReader(""))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	retstr = string(buf)
	return
}
func GetBetween(str, sstr, estr string) (rstr string, err error) {
	buf, sbuf, ebuf := []byte(str), []byte(sstr), []byte(estr)
	s := bytes.Index(buf, sbuf)
	if s == -1 {
		err = fmt.Errorf("未找到起始字符串:" + sstr)
		return
	}
	buf = buf[s+len(sbuf):]
	e := bytes.Index(buf, ebuf)
	if e == -1 {
		err = fmt.Errorf("未找到结束字符串:" + estr)
		return
	}
	buf = buf[:e]
	rstr = string(buf)
	return
}
func GetCurrentPath() (path string, err error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}
	i := strings.LastIndex(s, "\\")
	path = string(s[0 : i+1])
	return
}
