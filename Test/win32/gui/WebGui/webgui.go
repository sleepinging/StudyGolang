package main

import (
	"twt/mystr"
	"syscall"
	"net/http"
	"fmt"
	"log"
	"github.com/lxn/win"
	"time"
	"os"
)

func openbrowurl(url string) { //调用默认浏览器打开网页
	szOperation := mystr.Str2uft16ptr("open")
	szAddress := mystr.Str2uft16ptr(url)
	h, err := syscall.LoadLibrary("shell32.dll")
	if err != nil {
		abort("LoadLibrary", err)
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "ShellExecuteW")
	if err != nil {
		abort("GetProcAddress", err)
	}
	syscall.Syscall6(uintptr(proc), 6, 0, szOperation, szAddress, 0, 0, 1)
}
func abort(funcname string, err error) {
	panic(funcname + " failed: " + err.Error())
}

func LoginServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("收到请求")
	fmt.Println(req.PostFormValue("close"))
	name := req.PostFormValue("name")
	pwd := req.PostFormValue("password")
	fmt.Println("用户名", name, "密码", pwd)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200) //必须在所有header之后
	if name == pwd {
		fmt.Fprintf(w, `{"res":0,"err":""}`)
	} else {
		fmt.Fprintf(w, `{"res":-1,"err":""}`)
	}
}

func startserver() {
	http.HandleFunc("/login", LoginServer)
	fmt.Println("Server start...")
	err := http.ListenAndServe("127.0.0.1:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func findwindow(wname string) (f bool) {
	r := win.FindWindow(nil, mystr.Str2ui16p(wname))
	return r > 0
}

var isrun = false

func checkstate() {
	for !findwindow("Go界面") {
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("已经启动")
	isrun = true
	for findwindow("Go界面") {
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println("已经关闭")
	os.Exit(0)
}

func main() {
	go checkstate()
	//path,_:=mytools.GetCurrentPath()
	path := `cefsimple.exe`
	openbrowurl(path)
	startserver()
}
