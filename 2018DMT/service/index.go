package service

import (
	"net/http"
	"../global"
	"../tools"
	"../dao"
	"strings"
	"fmt"
	"time"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	if strings.HasSuffix(r.RequestURI, ".html") {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
			r.RemoteAddr, "访问", r.RequestURI)
	}
	global.RootFileServer.ServeHTTP(w, r)

	cookie, err := r.Cookie("user")
	go showinfo(cookie, r.RemoteAddr,err)
}

func showinfo(cookie *http.Cookie,ip string, err error) {
	if err != nil {
		return
	}
	user, t, err := dao.GetUserIdFromCookie(cookie.Value)
	fmt.Println(tools.FmtTime(), "Cookie userid:", user, t)
}