package router

import (
	"../../service"
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
	http.HandleFunc("/seekhelp/watchhelp", service.WatchSeekHelpHelps)
	//接受帮助
	http.HandleFunc("/seekhelp/accept", service.AcceptHelp)
	//回复帮助
	http.HandleFunc("/seekhelp/help/reply", service.ReplyHelp)
}

func RegisterAllRouter() {
	addLoginRouter()
	addJobRouter()
	addUserRouter()
	addCVRouter()
	addMsgRouter()
	addSeekHelpRouter()
	helpRouter()
	addOtherRouter()
}
