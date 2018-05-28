package main

import (
	"fmt"
	"twt/mytools"
	"net/http"
)

func main() {
	pt, _ := mytools.GetCurrentPath()
	http.Handle("/", http.FileServer(http.Dir(pt+`\wwwroot\`)))
	fmt.Println("服务启动...")
	http.ListenAndServe(":80", nil)
}
