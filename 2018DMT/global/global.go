package global

import (
	"twt/mytools"
	"net/http"
	"fmt"
)

var (
	CurrPath       string
	RootFileServer http.Handler
)

func init() {
	CurrPath, _ = mytools.GetCurrentPath()
	RootFileServer = http.FileServer(http.Dir(CurrPath + `\view\wwwroot\`))
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}