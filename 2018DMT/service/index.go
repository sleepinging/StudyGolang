package service

import (
	"net/http"
	"../models"
	"../global"
	"../tools"
	"../dao"
	"fmt"
	"strings"
	"time"
	"net/url"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	go func(r *http.Request){
		if strings.HasSuffix(r.RequestURI, ".html") {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"),
				r.RemoteAddr, "访问", r.RequestURI)
			dao.AddVisitRecord(time.Now(),r.RemoteAddr,r.RequestURI)
		}
	}(r)

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

func VisitRecordCount(w http.ResponseWriter, r *http.Request){
	queryForm, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	d,err:=GetGetInt("Day",queryForm)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	c,err:=dao.VisitRecordCount(d)
	if err != nil {
		models.SendRetJson2(0, "失败", err.Error(), w)
		return
	}
	models.SendRetJson2(1, "成功", c, w)
}