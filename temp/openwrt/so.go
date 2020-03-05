package main

/*
#include <malloc.h>
#include "loadso.h"
#cgo LDFLAGS: -ldl
*/
import "C"
import (
	"net"
	"fmt"
	"unsafe"
	"time"
)

var handle unsafe.Pointer
var fp unsafe.Pointer

func call(){
	srcmac:=net.HardwareAddr{1,2,3,4,5,6}
	dstmac:=net.HardwareAddr{6,5,4,3,2,1}
	srcip:=net.IPv4(192,168,0,1)
	dstip:=net.IPv4(10,0,0,1)
	srcport:=uint16(14567)
	dstport:=uint16(80)
	protocol:=uint8(6)
	appdata:=[]byte("GET /login?username=twt&pwd=123456 HTTP/1.1\r\nHost: www.baidu.com\r\n\r\n")
	tp:=0
	var id unsafe.Pointer//char*
	t1:=time.Now()
	f:=C.callfunc(C.FuncSign(fp),(*C.uchar)(unsafe.Pointer(&srcmac[0])),(*C.uchar)(unsafe.Pointer(&dstmac[0])),
		(*C.uchar)(unsafe.Pointer(&srcip[0])),(*C.uchar)(unsafe.Pointer(&dstip[0])),C.ushort(srcport),
		C.ushort(dstport),C.uchar(protocol),(*C.uchar)(unsafe.Pointer(&appdata[0])),C.uint(len(appdata)),
		(*C.int)(unsafe.Pointer(&tp)),(**C.char)(unsafe.Pointer(&id)))
	fmt.Println(C.int(f),C.int(tp),C.GoString((*C.char)(id)))
	C.free(id)
	fmt.Println("spent:",time.Now().Sub(t1))
}

func main() {
	handle=C.loadso(C.CString("/home/phei/wacapture/openwrt/plugins/tmptest.so"))
	fp=C.loadfunc(handle,C.CString("parse"))
	fpi:=C.loadfunc(handle,C.CString("init"))
	f:=C.callinit(C.FSinit(fpi))
	fmt.Println("init",C.int(f))
	for;;{
		call()
		time.Sleep(time.Millisecond*50)
	}
	C.unloadso(handle)
}
