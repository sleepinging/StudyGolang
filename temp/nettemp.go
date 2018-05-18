package main

import (
	"unsafe"
	"fmt"
)

type mybytes struct {
	p   uintptr
	len int
	cap int
}

func main() {
	b := []byte{1, 2, 3, 4, 5}
	pb := uintptr(unsafe.Pointer(&b))
	mb := (*mybytes)(unsafe.Pointer(pb))
	for i := 0; i < len(b); i++ {
		p := mb.p + uintptr(i)
		fmt.Print(*(*uint8)(unsafe.Pointer(p)))
	}
	pb += unsafe.Sizeof(mb.p)
	fmt.Println("len:", *(*int)(unsafe.Pointer(pb)))
	pb += unsafe.Sizeof(mb.len)
	fmt.Println("cap:", *(*int)(unsafe.Pointer(pb)))
}