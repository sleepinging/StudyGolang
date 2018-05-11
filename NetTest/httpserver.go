package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	rootpath = `E:/临时/wwwroot/`
	str404   = `
	<html>
	<head><title>未找到文件</title></head>
	<body bgcolor="white">
		<center><h1>未找到文件,请检查URL是否正确</h1></center>
	</body>
	</html>`
)

func readfile(filename string) (str string, err error) {
	_, err = os.Stat(filename)
	if err != nil {
		return
	}
	inputFile, fileerr := os.Open(filename)
	if fileerr != nil {
		return str, fileerr
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	bs := make([]byte, 1024*1024*1024*10) //10MB
	l := 0
	l, err = inputReader.Read(bs)
	if err != nil {
		return
	}
	str = string(bs[:l])
	return
}

var cc = 0

func DefaultServer(w http.ResponseWriter, req *http.Request) {
	cc++
	reqfile := rootpath + req.URL.Path[1:] //去除/
	fmt.Println(cc, ":", reqfile)
	if str, err := readfile(reqfile); err != nil {
		fmt.Println("404")
		fmt.Fprintf(w, str404)
	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, str)
	}
}

func startserver() {
	http.HandleFunc("/", DefaultServer)
	fmt.Println("Server start...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func readtest() {
	str, err := readfile("E:/临时/wwwroot/Rigster_test/index.jsp")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("内容:" + str)
}

func main() {
	//readtest()
	startserver()
}
