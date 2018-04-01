package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	purl "net/url"
)

func httpget(url string) {
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
}

func httppost(url string) {
	client := &http.Client{}
	var clusterinfo = purl.Values{}
	//var clusterinfo = map[string]string{}
	clusterinfo.Add("userName", "twt")
	data := clusterinfo.Encode()
	//提交请求
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		panic(err)
	}
	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	_, err = io.Copy(stdout, response.Body)
	//返回的状态码
	status := response.StatusCode

	fmt.Println(status)
}

func main() {
	url := "https://www.baidu.com"
	//httpget(url)
	httppost(url)
}
