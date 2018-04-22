package main

/*
#include <windows.h>
int myadd(int a,int b){
	return a+b;
}
int msgbox(wchar_t* c,wchar_t* t){
	return MessageBoxW(0,c,t,3);
}
*/
import "C"
import (
	"fmt"
	"twt/mystr"
	"unsafe"
)

func msgbox(text, title string) {
	r := C.msgbox((*C.wchar_t)(unsafe.Pointer(mystr.Str2uft16ptr(text))),
		(*C.wchar_t)(unsafe.Pointer(mystr.Str2uft16ptr(title))))
	res := C.int(r)
	fmt.Println(res)
}

func main() {
	r := C.myadd(1, 2)
	res := C.int(r)
	fmt.Println(res)
	msgbox("内容", "测试")
}
