package global

import "time"

var (
	MaxCookieTime = int((time.Hour * 10).Seconds()) //最大cookie时间
)
