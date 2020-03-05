package main

import (
	"fmt"
	"math/rand"
	"time"
	"twt/nettools"
	"sync"
)

var wg sync.WaitGroup

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
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklnmopqrstuvwxyz-+="
	bytes := []byte(str)
	result := make([]byte, plen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < plen; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

func GenNum(min,max int) (QQ string) {
	d := max - min
	q := rand.Intn(d) + min
	QQ = fmt.Sprintf("%d", q)
	return
}

//func GenIp() (IP string) {
//	IP = fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255)) + `.` +
//		fmt.Sprintf("%d", rand.Intn(255))
//	return
//}

func dd(){
	url:=`http://mcw22.com/home`
	for;;{
		//fmt.Println(nettools.HttpGet(url+"&qq=0"+GenQQ()))
		nettools.HttpGet(url+`?qq=`+GenNum(1,200))
	}
	wg.Done()
}

var c=0

func startfxxk(){
	headers := map[string]string{}
	headers["Referer"] = `http://jhj.qxd.bid/v0oxjT.php?_wv=89?oNgf`
	headers["Content-Type"] = `application/x-www-form-urlencoded`
	headers["User-Agent"] = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`
	url := `http://jhj.qxd.bid/r4lPaSKqaGje7VP.php`
	data:=``
	qq:=``
	pwd:=``
	wg:=&sync.WaitGroup{}
	for i:=0;i<999;i++{
		go func(){
			fmt.Println("1")
			wg.Add(1)
			for;;{
				c++
				qq=GenQQ()
				pwd=GenPwd()
				data=`tet_user=`+qq+`&tet_pass=`+pwd
				fmt.Println(c,qq,pwd)
				nettools.HttpPost(url,data,headers)
				//fmt.Println(nettools.HttpPost("http://svbe.massf.win/AE8J3DRU9K5Ph1VhDmkUO6gh1562nVA94g.php",data,headers))
			}
			wg.Done()
		}()

	}
	wg.Wait()
}

func main() {

	//for i:=0;i<999;i++{
	//	wg.Add(1)
	//	go dd()
	//}
	//fmt.Println("dd...")
	//wg.Wait()
	//fmt.Println(GenQQ(), GenPwd())
	//headers := map[string]string{}
	//headers["Referer"] = `http://jhj.qxd.bid/v0oxjT.php?_wv=89?oNgf`
	//headers["Content-Type"] = `application/x-www-form-urlencoded`
	//headers["User-Agent"] = `Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`
	//url := `http://jhj.qxd.bid/r4lPaSKqaGje7VP.php`
	//data:=`tet_user=123465456&tet_pass=1321323`
	//fmt.Println(nettools.HttpPost(url,data,headers))
	//startfxxk()
}