package main

import (
	"fmt"
	"twt/nettools"
)

func main() {
	url := "http://127.0.0.1:8765"
	h := make(map[string]string)
	h["1"] = "2"
	r, _ := nettools.HttpPost(url, "123456", h)
	fmt.Println(r)
}
