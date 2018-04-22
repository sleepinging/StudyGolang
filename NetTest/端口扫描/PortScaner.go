package main

import (
	"net"
	"fmt"
	"time"
	"math"
)

var (
	timeout = time.Millisecond * 100 //超时
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

func ScanPort(ip string, sport, eport int) { //扫描某IP的一段端口是否打开
	for port := sport; port < eport; port++ {
		if IsPortOpen(ip, port) {
			fmt.Println(ip, ":", port, "Open")
		}
	}
}

func main() {
	sport, eport := 10, 500
	maxth := 50
	ip := "115.239.210.27"
	plist := SplitPort(sport, eport, maxth)
	for i := 0; i < maxth; i++ {
		go ScanPort(ip, plist[i], plist[i+1])
	}
	fmt.Scanf("%d\n", &maxth)
}
