package main

import (
	"os"
	"bufio"
	"strings"
	"io"
	"crypto/des"
	"crypto/cipher"
	"encoding/base64"
	"sync"
	"path/filepath"
	"fmt"
	"flag"
)

var datach =make(chan []byte,100)
var wg =sync.WaitGroup{}
var savefile=`C:\Users\我啊\Desktop\bcp\raw.bcp`
var readpath=`C:\Users\我啊\Desktop\bcp`

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func DesDecrypt(crypted, key,iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func handlefile(filename string)(err error){
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func(){
		f.Close()
	}()
	rd := bufio.NewReader(f)
	c:=1
	for {
		c++
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		line=strings.TrimSuffix(line,"\n")
		if io.EOF == err{
			break
		}
		if len(line)==0{
			continue
		}
		//fmt.Println(line)
		data,err:=base64.StdEncoding.DecodeString(string([]byte(line)))
		if err != nil {
			fmt.Println("file:",filename," line:",c,":",line,err)
			continue
		}
		ndata,err:=DesDecrypt(data,[]byte("dx$@shdx"),[]byte("salt#&@!"))
		if err != nil {
			fmt.Println("file:",filename," line:",c,":",line,err)
			continue
		}
		//fmt.Print(string(ndata))
		datach<-ndata
		//fmt.Println(" down")
	}
	return
}

func appendfile(filename string, buf []byte) {
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()
	fd.Write(append(buf, '\n'))
}

func work(){
	wg.Add(1)
	for;;{
		data:=<-datach
		if len(data)==0{
			break
		}
		//appendfile(`C:\Users\我啊\Desktop\爱茉莉数据\sum\78a351229841-udp-raw.bcp`,append(data,'\n'))
		appendfile(savefile,data)
	}
	wg.Done()
}

func getFilelist(path string) {
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		fmt.Print(path)
		err=handlefile(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(" down!")
		//println(path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	datach<-[]byte{}
}

func main() {
	bp:=flag.String("path","","bcp files path")
	sp:=flag.String("save","","save file path")
	flag.Parse()
	savefile=*sp
	readpath=*bp
	if savefile==""||readpath==""{
		fmt.Println("please see --help")
		return
	}
	go work()
	getFilelist(readpath)
	wg.Wait()
}
