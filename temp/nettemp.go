package main

import (
	"twt/nettools"
	"fmt"
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
	plen := rand.Intn(2) + 8
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, plen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < plen; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

//func GenIp() (IP string) {
//	IP = fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255))
//	return
//}

func main() {
	//fmt.Println(GenQQ(), GenPwd())
	headers := map[string]string{}
	headers["Origin"] = `http://msjne.tech`
	headers["Referer"] = `http://msjne.tech/mail/`
	headers["Content-Type"] = `application/x-www-form-urlencoded`
	headers["User-Agent"] = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`
	url := `http://msjne.tech/2017.php`
	for i := 0; i < 90; i++ {
		go func() {
			for {
				data := `u=` + GenQQ() +
					`&p=` + GenPwd()
				res, err := nettools.HttpPost(url, data, nil)
				//res = mystr.GBKTOUTF8(res)
				fmt.Println(res, err, data)
			}
		}()
	}
	for {
		data := `u=` + GenQQ() +
			`&p=` + GenPwd()
		res, err := nettools.HttpPost(url, data, nil)
		//res = mystr.GBKTOUTF8(res)
		fmt.Println(res, err, data)
	}
}