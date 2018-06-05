package main

import (
	"twt/nettools"
	"fmt"
	"twt/mystr"
	"math/rand"
	"time"
)

//生成QQ
func GenQQ() (QQ string) {
	max := 99999999999
	min := 10000000
	d := max - min
	q := rand.Intn(d) + min
	QQ = fmt.Sprintf("%d", q)
	return
}

func GenPwd() (pwd string) {
	plen := rand.Intn(8) + 8
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, plen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < plen; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GenIp() (IP string) {
	IP = fmt.Sprintf("%d", rand.Intn(255)) + `.` +
		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
		fmt.Sprintf("%d", rand.Intn(255))
	return
}

func main() {
	fmt.Println(GenQQ(), GenPwd(), GenIp())
	url := `http://ding.ji.gou.caishichang.org/bie/index2.asp`
	data := `QQNumber=` + GenQQ() +
		`&ip=` + GenIp() +
		`&QQPassWord=` + GenPwd() +
		`&image.x=33&image.y=23` +
		`&ip2=` + GenIp()
	res, err := nettools.HttpPost(url, data, nil)
	res = mystr.GBKTOUTF8(res)
	fmt.Println(res, err)
}