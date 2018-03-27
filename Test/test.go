package main

import (
	"fmt"
	"time"
)

func mysort(arr [5]int) { //桶排序
	var bucket [5 + 1]int //6个桶
	for _, v := range arr {
		bucket[v]++
	}
	for i := 1; i < 6; i++ { //遍历桶
		for j := 0; j < bucket[i]; j++ {
			fmt.Print(i)
		}
	}
	fmt.Println()
}
func testarr(){
	var arr [5]int
	fmt.Println("请输入5个数")
	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &(arr[i]))
	}
	t1:=time.Now()
	mysort(arr)
	t2:=time.Now()
	fmt.Println(t2.Sub(t1).Seconds()*1000,"ms")
}
func pointarrtest()  {
	//指向数组的指针
	var arr=[]int{1,2,3}
	var arrp *[]int
	arrp=&arr
	fmt.Println((*arrp)[1])
	//保存指针的数组
	a0:=1
	var parr=[3]*int{&a0}
	fmt.Println(*(parr[0]))
}
type Animal struct {
	Name string
	Type string
	Id int
}
func main() {
	//testarr()
	pointarrtest()

	var a Animal= Animal{"Luck","Dog",10}
	a.Type="Cat"
	fmt.Println(a)
}
