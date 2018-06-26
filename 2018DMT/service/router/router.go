package router

import (
	"../../service"
	"../chart"
	"net/http"
)

func addLoginRouter() {
	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/login/islogin", service.IsLogin)
	http.HandleFunc("/login/logout", service.Logout)
}

func addJobRouter() {
	http.HandleFunc("/job/publish", service.PublishJob)
	http.HandleFunc("/job", service.ShowJob)
	http.HandleFunc("/job/query", service.QueryJob)
	http.HandleFunc("/job/querycount", service.QueryJobCount)
	http.HandleFunc("/job/updata", service.UpdataJob)
	http.HandleFunc("/job/delete", service.DeleteJob)
}

func addUserRouter() {
	http.HandleFunc("/user", service.GetUser)
	http.HandleFunc("/user/add", service.AddUser)
	http.HandleFunc("/user/updata", service.UpdateUser)
	http.HandleFunc("/user/delete", service.DeleteUser)

	//搜索用户
	http.HandleFunc("/user/search", service.SearchUser)

}

func addCVRouter() {
	http.HandleFunc("/cv/user", service.GetUserCV)
	http.HandleFunc("/cv", service.GetCV)
	http.HandleFunc("/cv/updata", service.UpDataCV)
	http.HandleFunc("/cv/delete", service.DeleteCV)
	http.HandleFunc("/cv/add", service.AddCV)
}

func addMsgRouter() {
	http.HandleFunc("/message", service.GetMsg)
	http.HandleFunc("/message/send", service.SendMsg)
	http.HandleFunc("/message/recved/count", service.RecvMsgCount)
	http.HandleFunc("/message/sended/count", service.SendMsgCount)
	http.HandleFunc("/message/sended", service.GetSendedMsg)
	http.HandleFunc("/message/recved", service.GetRecvedMsg)
	http.HandleFunc("/message/mark", service.MarkMsgRead)
	http.HandleFunc("/message/delete", service.DeleteMsg)
}

func addSeekHelpRouter() {
	http.HandleFunc("/seekhelp", service.GetSeekHelp)
	http.HandleFunc("/seekhelp/publish", service.PublishSeekHelp)
	http.HandleFunc("/seekhelp/search", service.SearchSeekHelp)
	http.HandleFunc("/seekhelp/updata", service.UpdataSeekHelp)
	http.HandleFunc("/seekhelp/delete", service.DeleteSeekHelp)
	http.HandleFunc("/seekhelp/search/count", service.CountSearchSeekhHelp)
}

func addOtherRouter() {
	http.HandleFunc("/", service.GetIndex)
	http.HandleFunc("/register", service.Register)
	http.HandleFunc("/register/sendcode", service.SendCode)
	http.HandleFunc("/upload/file", service.GetFileUrl)
	http.HandleFunc("/gold/show", service.GetUserGold)
	http.HandleFunc("/gold/set", service.SetUserGold)
}

func helpRouter() {
	//帮助求助
	http.HandleFunc("/seekhelp/help", service.PublishHelp)
	//获取某求助的帮助数
	http.HandleFunc("/seekhelp/help/count", service.GetHelpCount)
	//查看某求助的帮助
	http.HandleFunc("/seekhelp/gethelp", service.GetSeekHelpHelps)
	//接受帮助
	http.HandleFunc("/seekhelp/accept", service.AcceptHelp)
	//回复帮助
	http.HandleFunc("/seekhelp/help/reply", service.ReplyHelp)
	//获取某帮助的回复数量
	http.HandleFunc("/seekhelp/help/reply/count", service.CountHelpReply)
	//获取某帮助的回复
	http.HandleFunc("/seekhelp/help/getreply", service.GetHelpReply)
}

func chartRouter() {
	http.HandleFunc("/chart", chart.OnChartConnections)
	// Start listening for incoming chat messages
	go chart.OnChartMessages()
}

func blogRouter() {
	//发布博客
	http.HandleFunc("/blog/pulish", service.PublishBlog)
	//显示博客
	http.HandleFunc("/blog", service.GetBlog)
	//搜索博客
	http.HandleFunc("/blog/search", service.SearchBlog)
	//搜索博客的数量
	http.HandleFunc("/blog/search/count", service.CountSearchBlog)
	//修改博客
	http.HandleFunc("/blog/updata", service.UpdateBlog)
	//删除博客
	http.HandleFunc("/blog/delete", service.DeleteBlog)
	//点赞博客
	http.HandleFunc("/blog/zan", service.ZanBlog)
	//取消赞
	http.HandleFunc("/blog/zan/cancel", service.CancelZanBlog)
	//是否已赞
	http.HandleFunc("/blog/zan/check", service.CheckZanBlog)
	//回复博客
	http.HandleFunc("/blog/reply", service.ReplyBlog)
	//获取博客回复的数量
	http.HandleFunc("/blog/reply/count", service.CountBlogReply)
	//获取博客回复
	http.HandleFunc("/blog/reply/get", service.GetBlogReply)
	//查找某人发布的博客
	http.HandleFunc("/blog/search/user", service.GetUserBlogs)

	//查看某用户赞过的博客数
	http.HandleFunc("/user/zan/blog/count", service.CountUserZanBlogs)

	//查看某用户赞过的博客
	http.HandleFunc("/user/zan/blog", service.GetUserZanBlogs)

}

func statisticsRouter() {
	//访问量
	http.HandleFunc("/statistics/visit", service.VisitRecordCount)

	//注册量/用户数量
	http.HandleFunc("/statistics/register", service.RegisterRecordCount)

	//登录数量
	http.HandleFunc("/statistics/login", service.LoginRecordCount)

	//博客发布量
	http.HandleFunc("/statistics/blog/publish", service.BlogPublishCount)

	//求助发布量
	http.HandleFunc("/statistics/seekhelp/publish", service.SeekHelpPublishCount)

	//工作发布量
	http.HandleFunc("/statistics/job/publish", service.JobPublishCount)

}

func RegisterAllRouter() {
	addLoginRouter()
	addJobRouter()
	addUserRouter()
	addCVRouter()
	addMsgRouter()
	addSeekHelpRouter()
	helpRouter()
	chartRouter()
	blogRouter()
	statisticsRouter()
	addOtherRouter()
}
