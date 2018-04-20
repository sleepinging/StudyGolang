package main

import (
	"twt/mystr"
	"syscall"
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
func msgbox3(text, title string) (res int32) {
	h, err := syscall.LoadLibrary("user32.dll")
	if err != nil {
		abort("LoadLibrary", err)
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "MessageBoxW")
	if err != nil {
		abort("GetProcAddress", err)
	}
	r, _, _ := syscall.Syscall6(uintptr(proc), 4, 0, mystr.Str2uft16ptr(text), mystr.Str2uft16ptr(title), 3, 0, 0)
	res = int32(r)
	return
}
func main() {
	//openbrowurl("www.baidu.com")
	c, t := "你好", "提示"
	pc, pt := mystr.Str2uft16ptr(c), mystr.Str2uft16ptr(t)
	c, t = mystr.Utf16prt2str(pc), mystr.Utf16prt2str(pt)
	msgbox3(c, t)
}
