package main

import (
	"./tools"
	"fmt"
	"math"
)

var (
	stopf = false
)

func splitpwd(sxh, exh int64, num int) (plist []int64) { //将学号均匀分割
	nd := exh - sxh
	_ = nd
	plist = make([]int64, num+1)
	plist[0] = sxh
	for i := 1; i < num; i++ {
		a := math.Floor(float64(nd) / float64(num) * float64(i))
		p := int64(a) + sxh
		//fmt.Println(p)
		plist[i] = p
	}
	plist[num] = exh
	return
}

func startsearch(sxh, exh int64) {
	for xh := sxh; xh < exh && !stopf; xh++ {
		name, err := tools.GetStudentName(xh)
		if err != nil {
			fmt.Println("未找到学号", xh)
			continue
		}
		fmt.Printf("%c[1;40;32m%s%s%s%d%c[0m", 0x1B, "找到姓名:", name, "学号:", xh, 0x1B)
		go tools.InsertInfo(xh, name)
	}
}

func bianlixh() {
	maxthread := 3
	//fmt.Println("输入线程数:")
	//fmt.Scan(&maxthread)
	//xhlist:=splitpwd(201454000000,201799999999,maxthread)
	//for i:=0;i<maxthread;i++{
	//	go startsearch(xhlist[i],xhlist[i+1])
	//	fmt.Println("线程",i,"启动成功")
	//	//fmt.Println(xhlist[i],xhlist[i+1])
	//}
	fmt.Println("按回车键结束")
	startsearch(201454000000, 201799999999)
	fmt.Scanf("%d\n", &maxthread)
	stopf = true
	fmt.Println("已经停止遍历")
}
func getname() {
	xh := int64(0)
	fmt.Println("输入学号:")
	fmt.Scan(&xh)
	fmt.Println("查询中...")
	name, err := tools.GetStudentName(xh)
	if err != nil {
		fmt.Println("未找到！请重新输入学号")
		return
	}
	fmt.Println(xh, ":"+name)
}
func menu() {
LABEL_MENU:
	fmt.Println("选择功能:")
	fmt.Println("1.根据学号查找姓名")
	fmt.Println("2.遍历学号")
	fmt.Println("3.退出")
	gn := 0
	_, err := fmt.Scan(&gn)
	if err != nil || gn < 1 || gn > 3 {
		goto LABEL_MENU
	}
	switch gn {
	case 1:
		getname()
	case 2:
		bianlixh()
	case 3:
		return
	default:

	}
	goto LABEL_MENU
}
func main() {
	menu()
}
