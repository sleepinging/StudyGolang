package main

import (
	"os"
	"fmt"
)

func main() {
	r,w,err:=os.Pipe()
	if err != nil {
		fmt.Println(err)
	}
	i:=0
	fmt.Scanf("%d",&i)
	fmt.Println(w.Name())
	fmt.Println(r.Name())
	w.WriteString("123")
	var s = make([]byte, 20)
	n, _ := r.Read(s)
	fmt.Println(string(s[:n]))
	_,_=w,r
}
