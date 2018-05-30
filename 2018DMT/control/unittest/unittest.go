package unittest

import "fmt"
import "../EmailVerify"

func TestEmailVerify() {
	email := "237731947@qq.com"
	code := EmailVerify.GenCode()
	fmt.Println(code)
	EmailVerify.UpdateCode(email, code)
	//EmailVerify.SendCode(email)
	fmt.Scanln(&code)
	//fmt.Println(code)
	fmt.Println(EmailVerify.CheckCode(email, code))
}

func Test() {
	//TestEmailVerify()
}
