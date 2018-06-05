package dao

import (
	"time"
	"../global"
	"../tools"
	"strings"
	"strconv"
)

func GenUserCookie(user string) (cookie string) {
	t := time.Now().Format("2006-01-02/15:04:05")
	cookie, _ = tools.Encrypt(user+` `+t, global.MiKey)
	return
}

func GetUserIdFromCookie(cookie string) (id int, t time.Time, err error) {
	dec, _ := tools.Decrypt(cookie, global.MiKey)
	us := strings.Split(dec, ` `)
	if len(us) < 1 {
		return
	}
	uid := us[0]
	t, err = time.Parse("2006-01-02/15:04:05", us[1])
	if err != nil {
		return
	}
	if uid == "" {
		err = global.NotLogin
		return
	}
	id64, err := strconv.ParseInt(uid, 10, 32)
	if err != nil {
		return
	}
	id = int(id64)
	return
}
