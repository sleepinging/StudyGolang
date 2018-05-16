package main

//
import "C"
import (
	"encoding/json"
	"fmt"
	"github.com/axgle/mahonia"
	"unsafe"
	"net/http"
	"strings"
	"io/ioutil"
)

//export Sum
func Sum(a int, b int) int { //最简单的计算和
	return a + b
}

//export Show
func Show(str string) { //显示传入的字符串
	fmt.Print(str)
}

//export ToNewGBKCStr
func ToNewGBKCStr(str string, c **C.char) { //生成gbk字符串
	enc := mahonia.NewEncoder("GBK")
	output := enc.ConvertString(str)
	*c = C.CString(output)
}

//export HttpGet
func HttpGet(url string, c **C.char) { //http请求
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, strings.NewReader(""))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		fmt.Println(err)
		return
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	retstr := string(buf)
	ToNewGBKCStr(retstr, c)
}

//export SetMyPoint
func SetMyPoint(p uintptr) { //修改传入的结构体
	mp := (*MyPoint)(unsafe.Pointer(p)) //将指针p强制转换为MyPoint类型
	mp.Name = "名称:" + mp.Name
	mp.X = 1 + mp.X
	mp.Y = 2 + mp.Y
}

//export JsonToMyPoint
func JsonToMyPoint(str string, p uintptr) {
	mp := (*MyPoint)(unsafe.Pointer(p))
	json.Unmarshal([]byte(str), mp)
}

type MyPoint struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

func main() {
	//测试golang结构体内存布局
	mp := MyPoint{Name: "测试", X: 1, Y: 2}
	p := uintptr(unsafe.Pointer(&mp))     //强制转换为指针
	name := *(*string)(unsafe.Pointer(p)) //将指针强制转换为string
	fmt.Println(name)

	//64位系统上string占16字节，前8字节是字符串的起始地址，后8字节为字符串占的字节数
	p += unsafe.Sizeof(string(""))
	x := *(*int)(unsafe.Pointer(p)) //将指针强制转换为int
	fmt.Println(x)

	p += unsafe.Sizeof(int(0)) //64位系统上int占8字节
	y := *(*int)(unsafe.Pointer(p))
	fmt.Println(y)
}
