package main

import (
	"github.com/slclub/goevents"
	"fmt"
)

var ev1 = func(s1 string,a int,s2 string){
	fmt.Println(s1,a,s2)
}
var ev2 = func(s string, a1,a2 int){
	fmt.Println(s,a1,a2)
}

func main() {
	bx()
}


//串行事件
func cx(){
	events := goevents.Classic()
	//连贯写法 绑定事件及参数
	events.Bind("123", "dfd").On("a", ev1)

	//绑定事件和参数分开写
	events.Bind("event2 serial running")
	events.On("a",ev2)
	//Trigger
	//events.Bind(...)//可以在这里绑定全局参数，当有事件绑定参数，这里可以重新绑定，只是不能针对所有串行事件
	//按模块名 执行事件队列。无参数 执行所有串行 事件
	events.Trigger("a")
	//触发所有事件，串行，并行的全部触发
	//events.Emit()
}

//并行事件
func bx(){
	events := goevents.Classic()

	//绑定并发事件
	events.GoOn(ev1, "123",344,"event1 param")
	events.GoOn(ev2, "event2 param",454, 343)
	//最终执行事件。
	//events.End(func(...interface{}){}, "123","End 事件 最后所有事件执行完 才会执行")
	//触发并发事件不能用Trigger
	events.Emit()
}