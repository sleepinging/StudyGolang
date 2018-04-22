package main

import (
	"syscall"
	"unsafe"
	"unicode/utf16"
)

func main() {
	h, err := syscall.LoadLibrary("user32.dll")
	if err != nil {
		abort("LoadLibrary", err)
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "MessageBoxW")
	if err != nil {
		abort("GetProcAddress", err)
	}
	text, title := "内容", "标题"
	syscall.Syscall6(uintptr(proc), 4, 0, str2uiptr(text), str2uiptr(title), 0, 0, 0)
}

func str2uiptr(str string) (p uintptr) { //将字符串转为uft16指针
	e := utf16.Encode([]rune(str))     //转成unicode
	e = append(e, uint16(0))           //添加末尾的0
	p = uintptr(unsafe.Pointer(&e[0])) //转成指针
	//p=uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(str)))
	return
}

func abort(funcname string, err error) {
	panic(funcname + " failed: " + err.Error())
}
