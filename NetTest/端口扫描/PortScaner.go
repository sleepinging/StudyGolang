package main

import (
	"net"
	"fmt"
	"time"
	"math"
)

var (
	timeout  = time.Millisecond * 100 //超时
	stopflag = false                  //用于停止所有扫描线程
)

func SplitPort(sport, eport int, num int) (plist []int) { //将端口号均匀分割
	nd := eport - sport
	_ = nd
	plist = make([]int, num+1)
	plist[0] = sport
	for i := 1; i < num; i++ {
		a := math.Floor(float64(nd) / float64(num) * float64(i))
		p := int(a) + sport
		plist[i] = p
	}
	plist[num] = eport
	return
}

func IsPortOpen(ip string, port int) (isopen bool) { //检测某IP的某个端口是否打开
	_, err := net.DialTimeout("tcp", ip+":"+fmt.Sprintf("%d", port), timeout)
	return err == nil
}

func ScanPort(ip string, sport, eport int, portch chan int) { //扫描某IP的一段端口
	go exportport(ip, portch)
	for port := sport; port < eport && !stopflag; port++ {
		if IsPortOpen(ip, port) {
			//fmt.Println(ip, ":", port, "Open")
			portch <- port
		}
	}
}

func exportport(ip string, portch <-chan int) {
	for {
		fmt.Println(ip+":", <-portch)
	}
}

func StartScan(ip string, sport, eport int) {
	maxth := 50
	plist := SplitPort(sport, eport, maxth)
	for i := 0; i < maxth; i++ {
		portch := make(chan int, 5)
		go ScanPort(ip, plist[i], plist[i+1], portch)
	}
	fmt.Println("按回车键停止")
	fmt.Scanf("%d\n", &maxth)
	stopflag = true
}

func main() {
	sport, eport := 10, 500
	ip := "115.239.210.27"
	StartScan(ip, sport, eport)
}
