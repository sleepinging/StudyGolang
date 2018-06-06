package tools

import (
	"net/http"
	"strings"
	"io/ioutil"
)

func HttpGet(url string) (retstr string, err error) {
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, err := http.NewRequest("GET", url, strings.NewReader(""))
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	retstr = string(buf)
	return
}

func HttpPost(url string, data string, headers map[string]string) (retstr string, err error) { //发送post
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, strings.NewReader(data))
	if err != nil {
		//fmt.Println("发送失败!")
		return
		//panic(err)
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		reqest.Header.Set(k, v)
	}
	//处理返回结果
	response, err := client.Do(reqest)
	if err != nil {
		return
	}
	buf, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	retstr = string(buf)
	return
}
