package main

/*
#include <windows.h>
int myadd(int a,int b){
	return a+b;
}
void msgbox(wchar_t* c,wchar_t* t){
	MessageBoxW(0,c,t,0);
}
*/
import "C"
import (
	"fmt"
	"twt/mystr"
	"unsafe"
)

func msgbox(text, title string) {
	C.msgbox((*C.wchar_t)(unsafe.Pointer(mystr.Str2uft16ptr(text))),
		(*C.wchar_t)(unsafe.Pointer(mystr.Str2uft16ptr(title))))
}

func main() {
	r := C.myadd(1, 2)
	fmt.Println(r)
	msgbox("内容", "测试")
}
