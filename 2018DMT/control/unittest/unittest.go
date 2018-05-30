package unittest

import "fmt"
import (
	"../EmailVerify"
	"../../control"
	"net/http"
	"time"
	"twt/mytools"
	"twt/nettools"
)

func TestEmailVerify() {
	email := "1449693643@qq.com"
	code := EmailVerify.GenCode()
	fmt.Println(code)
	//EmailVerify.UpdateCode(email, code)
	EmailVerify.SendCode(email)
	fmt.Scanln(&code)
	//fmt.Println(code)
	fmt.Println(EmailVerify.CheckCode(email, code))
}

func getroot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), r.RemoteAddr, "访问")
	pt, _ := mytools.GetCurrentPath()
	fs := http.FileServer(http.Dir(pt + `\view\wwwroot\`))
	fs.ServeHTTP(w, r)
}

func webtest() {
	http.HandleFunc("/", getroot)
}

func PostTest() {
	url := "http://193.112.77.180/register/sendcode"
	data := "Email=237731947@qq.com"
	res, err := nettools.HttpPost(url, data, nil)
	control.CheckErr(err)
	fmt.Println(res)
}

func Test() {
	PostTest()
	//webtest()
	//TestEmailVerify()
}
