package global

import "errors"

var (
	LoginCookiePass = errors.New("登录过期")
	NoSuchJob       = errors.New("没有这个工作")
	NoSuchUser      = errors.New("没有这个用户")
	NoSuchCV        = errors.New("没有这个简历")
	UserNoCV        = errors.New("该用户没有简历")
	NoSuchBlog      = errors.New("没有这个博客")
	NoSuchHelp      = errors.New("没有这个帮助")
	NoSuchMsg       = errors.New("没有这个消息")
	EmptyUserPwd    = errors.New("用户名或密码不能为空")
	NotLogin        = errors.New("用户未登录")
	NoPermission    = errors.New("没有权限")
	NotYourMsg      = errors.New("不是你的消息")
)
