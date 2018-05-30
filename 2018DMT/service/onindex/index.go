package onindex

import (
	"net/http"
	"fmt"
	"time"
	"../../models/global"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println(time.Now(), r.RemoteAddr, "访问", r.RequestURI)
	global.RootFileServer.ServeHTTP(w, r)
}
