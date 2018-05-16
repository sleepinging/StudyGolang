package main

//
import "C"
import (
	"fmt"
	"unsafe"
	"github.com/axgle/mahonia"
)

//export Sum
func Sum(a int, b int) int {
	return a + b
}

//export Show
func Show(str string) {
	fmt.Print(str)
}

//export ToNewCStr
func ToNewCStr(str string, c **C.char) {
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(str)
	*c = C.CString(output)
}

//export SetMyPoint
func SetMyPoint(p uintptr) {
	mp := (*MyPoint)(unsafe.Pointer(p))
	mp.Name = "测试传递结构体"
	mp.X = 1
	mp.Y = 2
}

type MyPoint struct {
	Name string `名字`
	X    int    `x坐标`
	Y    int    `y坐标`
}

func main() {
	mp := MyPoint{Name: "测试", X: 1, Y: 2}
	p := uintptr(unsafe.Pointer(&mp))
	name := *(*string)(unsafe.Pointer(p))
	fmt.Println(name)

	p += unsafe.Sizeof(string(""))
	x := *(*int)(unsafe.Pointer(p))
	fmt.Println(x)

	p += unsafe.Sizeof(int(0))
	y := *(*int)(unsafe.Pointer(p))
	fmt.Println(y)
}
