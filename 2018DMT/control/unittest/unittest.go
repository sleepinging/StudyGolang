package unittest

import (
	"fmt"
	"../EmailVerify"
	"../../global"
	"../../models"
	"../../dao"
	"net/http"
	"time"
	"twt/mytools"
	"twt/nettools"
	"encoding/json"
	"io/ioutil"
	"os"
)

func TestEmailVerify() {
	email := "237731947@qq.com"
	//EmailVerify.UpdateCode(email, code)
	code := EmailVerify.SendCode(email)
	fmt.Println(code)
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
	global.CheckErr(err)
	fmt.Println(res)
}

func LoginTest() {
	e := dao.CheckLogin(&models.Login{Email: "237731947@qq.com", Password: "123456"})
	fmt.Println(e)
}

func jsonTest() {
	str := `{"email":"23@2","password":"123"}`
	o := &models.Login{}
	err := json.Unmarshal([]byte(str), o)
	global.CheckErr(err)
	fmt.Println(o)
}

func PostLoginTest() {
	time.Sleep(time.Second * 3)
	url := "http://127.0.0.1/login"
	data := "Email=237731947@qq.com&Password=123456"
	res, err := nettools.HttpPost(url, data, nil)
	global.CheckErr(err)
	fmt.Println(res)
}

func ConfigTest() {
	file, _ := os.OpenFile(global.CurrPath+`config.json`, os.O_RDONLY, os.ModeType)
	bs, _ := ioutil.ReadAll(file)
	cfg := models.Config{}
	json.Unmarshal(bs, &cfg)
	fmt.Println(cfg)
}

func Test() {
	//ConfigTest()
	//PostLoginTest()
	//jsonTest()
	//LoginTest()
	//PostTest()
	//webtest()
	//TestEmailVerify()
}
