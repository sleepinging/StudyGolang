package global

import (
	"twt/mytools"
	"net/http"
)

var (
	CurrPath       string
	RootFileServer http.Handler
)

func init() {
	CurrPath, _ = mytools.GetCurrentPath()
	RootFileServer = http.FileServer(http.Dir(CurrPath + `\view\wwwroot\`))
}
