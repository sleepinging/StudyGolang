package main

import (
	"encoding/json"
	"fmt"
)

type MyRecord struct {
	Person Person `json:"person"`
	Time   string `json:"time"`
	Place  string `json:"place"`
	Device int    `json:"device"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var jstr = `{
  "person": {
    "name": "小明",
    "age": 15
  },
  "time": "2018-01-01T15:00",
  "place": "5",
  "device": 10
}`

func main() {
	var o MyRecord
	err := json.Unmarshal([]byte(jstr), &o)
	if err != nil {
		fmt.Printf("解析错误%v", err)
		return
	}
	fmt.Println(o.Person.Name)
	o.Person.Name = "小红"
	b, err := json.Marshal(&o)
	if err != nil {
		fmt.Printf("err was %v", err)
		return
	}
	fmt.Println(string(b))
}
