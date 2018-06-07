package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"twt/mystr"
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
func GetCurrentPath() (path string, err error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return
	}
	i := strings.LastIndex(s, "\\")
	path = string(s[0 : i+1])
	return
}
func main() {
	path, _ := GetCurrentPath()
	fmt.Println(path)
}
