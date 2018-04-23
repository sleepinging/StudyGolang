package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"sync"
	"time"
)

var (
	timeout = time.Millisecond * 100 //超时
	//stopflag        = false                  //用于停止所有扫描线程
	maxth           = 5        //最大线程
	completedth     = 0        //已完成
	allportnum      = 0        //所有要扫描的端口数
	completeportnum = 0        //已完成扫描的端口
	nummux          sync.Mutex //是否在读写端口数加锁
	thmux           sync.Mutex //是否在读写线程数加锁
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
	for port := sport; port < eport; /*&& !stopflag*/ port++ {
		if IsPortOpen(ip, port) {
			//fmt.Println(ip, ":", port, "Open")
			portch <- port
		}
		go updateprocess()
		//fmt.Println(float64(completeportnum)/float64(allportnum)*100,"%")
	}
	go updatethreadnum()
}

func updatethreadnum() {
	thmux.Lock()
	completedth++
	thmux.Unlock()
}

func updateprocess() { //更新进度
	nummux.Lock()
	completeportnum++
	if completeportnum == allportnum {
		fmt.Print("\r扫描完毕")
	} else {
		str := fmt.Sprintf("%.1f", float64(completeportnum)/float64(allportnum)*100)
		fmt.Print("\r" + str + "%")
	}
	nummux.Unlock()
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

func showinfo() {
	fmt.Println("本软件开源，项目地址：")
	fmt.Println("https://github.com/sleepinging/StudyGolang/tree/master/NetTest/%E7%AB%AF%E5%8F%A3%E6%89%AB%E6%8F%8F")
	fmt.Println("使用 -h命令查看帮助")
}

func main() {
	showinfo()
	ip := flag.String("ip", "127.0.0.1", "你想要扫描的IP")
	sport := flag.Int("sp", 0, "起始端口")
	eport := flag.Int("ep", 1023, "结束端口")
	imaxth := flag.Int("mt", 5, "最大线程数")
	flag.Parse()
	maxth = *imaxth
	ns, err := net.LookupHost(*ip)
	if err != nil {
		fmt.Println("IP或域名输入有误", err)
		return
	}
	*ip = ns[0]
	if *eport < *sport || *eport > 65535 || *sport < 0 {
		fmt.Println("要求结束端口不小于起始端口且在0-65535之间")
		return
	}
	if maxth < 1 {
		fmt.Println("最大线程数应该大于0")
		return
	}
	fmt.Println("扫描", *ip+":", *sport, "-", *eport, "线程数:", maxth)
	allportnum = *eport - *sport + 1
	StartScan(*ip, *sport, *eport+1, maxth)
}
