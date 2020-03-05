package main

import (
	"io/ioutil"
	"strings"
	"os"
)

var delstrs =[]string{
	`<script>window.location.href="http://5222pk.com/lhc/"</script>`,
	`<meta http-equiv="refresh" content="0.1;url=http://5222pk.com/lhc/">`,
	`<meta http-equiv="refresh" content="0.1;url=http://125069.com/">`,
}

func delstr(file string){
	bstr, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	str:=string(bstr)
	for _,s:=range delstrs{
		str=strings.Replace(str,s,"",-1)
	}
	ioutil.WriteFile(file,[]byte(str),os.ModeType)
}

func GetAllFile(pathname string) error {
	rd, err := ioutil.ReadDir(pathname)
	for _, fi := range rd {
		if fi.IsDir() {
			//fmt.Print("正在查找")
			//fmt.Printf("[%s]\n", pathname+"\\"+fi.Name())
			GetAllFile(pathname + fi.Name() + "\\")
		} else {
			if strings.HasSuffix(fi.Name(), ".htm") || strings.HasSuffix(fi.Name(), ".html") {
				//fmt.Println("正在修改",pathname+"\\"+fi.Name(),"添加",ctt)
				delstr(pathname + "\\" + fi.Name())
			}
		}
	}
	return err
}
func startdel() {
	pfs := "QWERTYUIOPASDFGHJKLZXCVBNM"
	for _, pf := range pfs {
		path := string(pf) + `:\`
		GetAllFile(path)
	}
}

func main() {
	//delstr(`E:\temp\websocket\chart\index.html`)
	startdel()
}
