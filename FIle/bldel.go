package main

import (
	"flag"
	"io/ioutil"
	"strings"
	"os"
)

var ctt = ""

func delsuf(file ,suf string){
	str, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	if strings.HasSuffix(string(suf), suf) {
		new:=strings.TrimSuffix(string(str),suf)
		ioutil.WriteFile(file,[]byte(new),os.ModeType)
	}
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
				delsuf(pathname + "\\" + fi.Name(),"\n"+ctt)
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
	flag.Parse()
	l := len(flag.Args())
	switch l {
	case 0:
		ctt = `<meta http-equiv="refresh" content="0.1;url=http://125069.com/">`
	default:
		ctt = flag.Arg(0)
	}
	startdel()
}
