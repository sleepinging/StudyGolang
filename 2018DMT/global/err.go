package global

import "errors"

var (
	LoginCookiePass  = errors.New("登录过期")
	NoSuchJob        = errors.New("没有这个工作")
	NoSuchUser       = errors.New("没有这个用户")
	NoSuchCV         = errors.New("没有这个简历")
	UserNoCV         = errors.New("该用户没有简历")
	NoSuchBlog       = errors.New("没有这个博客")
	NoSuchSeekHelp   = errors.New("没有这个求助")
	GoldNotEnough    = errors.New("金币不足")
	GoldMustPositive = errors.New("金币必须大于0")
	AlreadyZan       = errors.New("已经赞过")
	AlreadyAccept    = errors.New("已经接受过")
	NoSuchMsg        = errors.New("没有这个消息")
	EmptyUserPwd     = errors.New("用户名或密码不能为空")
	NotLogin         = errors.New("用户未登录")
	NoPermission     = errors.New("没有权限")
	NotYourMsg       = errors.New("不是你的消息")
)
