package dao

import (
	"time"
	"../global"
	"../tools"
	"strings"
)

func GenUserCookie(user string) (cookie string) {
	t := time.Now().Format("2006-01-02/15:04:05")
	cookie, _ = tools.Encrypt(user+` `+t, global.MiKey)
	return
}

func GetUserinfo(cookie string) (user string, t time.Time, err error) {
	dec, _ := tools.Decrypt(cookie, global.MiKey)
	us := strings.Split(dec, ` `)
	if len(us) < 1 {
		return
	}
	user = us[0]
	t, err = time.Parse("2006-01-02/15:04:05", us[1])
	if err != nil {
		return
	}
	return
}
