package main

import (
	"twt/mystr"
	"syscall"
	"net/http"
	"fmt"
	"log"
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
	name := req.PostFormValue("name")
	pwd := req.PostFormValue("pwd")
	fmt.Println(name, pwd)
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200) //必须在所有header之后
	fmt.Fprintf(w, `{"res":0,"err":""}`)
}

func startserver() {
	http.HandleFunc("/login", LoginServer)
	fmt.Println("Server start...")
	err := http.ListenAndServe("127.0.0.1:8765", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func main() {
	//path,_:=mytools.GetCurrentPath()
	path := `E:\软件\学习中\Go\Study\StudyGolang\Test\win32\gui\WebGui\`
	openbrowurl(path + `1.html`)
	startserver()
}
