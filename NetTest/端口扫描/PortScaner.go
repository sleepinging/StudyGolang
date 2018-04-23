package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"time"
)

var (
	timeout         = time.Millisecond * 100 //超时
	stopflag        = false                  //用于停止所有扫描线程
	maxth           = 5                      //最大线程
	completedth     = 0                      //已完成
	allportnum      = 0                      //所有要扫描的端口数
	completeportnum = 0                      //已完成扫描的端口
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
		completeportnum++
		//fmt.Println(float64(completeportnum)/float64(allportnum)*100,"%")
	}
	completedth++
}

func exportport(ip string, portch <-chan int) {
	for {
		fmt.Println(ip+":", <-portch)
	}
}

func StartScan(ip string, sport, eport, maxth int) {
	plist := SplitPort(sport, eport, maxth)
	for i := 0; i < maxth; i++ {
		portch := make(chan int, 5)
		go ScanPort(ip, plist[i], plist[i+1], portch)
	}
	//fmt.Println("按回车键停止")
	for completedth < maxth {
		time.Sleep(time.Millisecond * 200)
	}
	//fmt.Scanf("%d\n", &maxth)
	//stopflag = true
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "你想要扫描的IP")
	sport := flag.Int("sp", 0, "起始端口")
	eport := flag.Int("ep", 65535, "结束端口")
	imaxth := flag.Int("mt", 5, "最大线程数")
	flag.Parse()
	maxth = *imaxth
	fmt.Println("扫描", *ip+":", *sport, "-", *eport, "线程数:", maxth)
	allportnum = *eport - *sport + 1
	StartScan(*ip, *sport, *eport+1, maxth)
	//sport, eport := 10, 500
	//ip := "115.239.210.27"
	//StartScan(ip, sport, eport,50)
}
