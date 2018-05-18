package main

import (
	"github.com/dchest/captcha"
	"bytes"
	"io/ioutil"
	"os"
	"fmt"
)

func simple() {
	id := captcha.New()
	buf := bytes.NewBuffer(make([]byte, 0))
	captcha.WriteImage(buf, id, captcha.StdWidth, captcha.StdHeight)
	ioutil.WriteFile("E:/1.jpg", buf.Bytes(), os.ModeAppend)

	v := "请输入验证码"
	fmt.Println(v)
	fmt.Scanln(&v)
	f := captcha.VerifyString(id, v)
	fmt.Println(f)
}

func main() {
	simple()
}
