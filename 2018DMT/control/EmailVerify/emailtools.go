package EmailVerify

import (
	"net/smtp"
	"strings"
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

func SendEmail(to, subject, text string) (err error) {
	user := "verify@beyondsky.club"
	password := "123456"
	host := "smtp.ym.163.com:25"
	//user := "verify@beyondsky.club"
	//password := "123456"
	//host := "smtp.ym.163.com:25"
	body := `
		<html>
		<body>
		<h3>
		` + text + `
		</h3>
		</body>
		</html>
		`
	//fmt.Println("send Email")
	err = sendToMail(user, password, host, to, subject, body, "html")
	return
}
