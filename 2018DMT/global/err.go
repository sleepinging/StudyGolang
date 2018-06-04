package global

import "errors"

var (
	LoginCookiePass = errors.New("登录过期")
	NoSuchJob       = errors.New("没有这个工作")
	NoSuchUser      = errors.New("没有这个用户")
)
