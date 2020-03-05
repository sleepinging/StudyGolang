package main

import (
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/sciter-sdk/go-sciter"
	"time"
	"log"
)

func main() {
	//extradll()
	rect := sciter.NewRect(400, 600, 300, 200)
	w, err := window.New(sciter.SW_TITLEBAR| sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|sciter.SW_MAIN|sciter.SW_ENABLE_DEBUG|sciter.SW_OWNS_VM,
		rect)
	if err != nil {
		log.Fatal(err)
	}
	//os.Chdir("view/login/mylogin")
	w.LoadFile("view/login.html")
	//w.LoadHtml("", "")

	w.SetTitle("青少年法律知识试题")
	defFunc(w)
	w.Show()
	w.Run()
}

func updatetime(w *window.Window){
	for {
		w.Call("settime",sciter.NewValue(time.Now().Format("15:04:05")))
		time.Sleep(time.Second)
	}
}

func login(name ,pwd string,w *window.Window)(res string){
	res="登录失败"
	//if name==`admin`&&pwd==`admin`{
	if name==`admin`&&pwd==`admin`{
		res="登录成功"
		w.LoadFile("view/su.html")
		w.Call("rsz",sciter.NullValue())
		w.Show()
	}
	return
}

func defFunc(w *window.Window) {
	w.DefineFunction("login", func(args ...*sciter.Value) *sciter.Value {
		name,pwd:=args[0].String(),args[1].String()
		//fmt.Println(name,pwd)
		return sciter.NewValue(login(name,pwd,w))
		//return sciter.NullValue()
	})
	go updatetime(w)
}
