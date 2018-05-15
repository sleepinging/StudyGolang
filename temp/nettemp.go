package main

import (
	"fmt"
	"twt/nettools"
	"strings"
	"encoding/json"
)

type jsret struct {
	Status int    `json:"status"`
	List   string `json:"list"`
}

func GetShort() (short string) {
	url := `http://www.ft12.com/create.php?m=index&a=urlCreate`
	res, err := nettools.HttpPost(url,
		"url=123.207.91.166%2Flove.html&type=6",
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	res = strings.Replace(res, `\`, ``, -1)
	jso := new(jsret)
	json.Unmarshal([]byte(res), jso)
	res = jso.List
	//fmt.Println(res)
	short = res[len(res)-5:]
	//fmt.Println(res)
	return
}

func main() {
	for res := GetShort(); ; res = GetShort() {
		fmt.Println(res)
		if res == `aAtwt` {
			break
		}
		//time.Sleep(time.Millisecond*300)
	}
}
