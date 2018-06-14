package main

import (
	"fmt"
	"github.com/pysrc/bs"
	"os"
	"strings"
	"twt/mystr"
	"twt/nettools"
	"sync"
)

var (
	Host = "http://www.biqukan.com"
	thch=make(chan bool,50)
	wg=sync.WaitGroup{}
)

//type sstrs []string
//func (s sstrs) Len() int           { return len(s) }
//func (s sstrs) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
//func (s sstrs) Less(i, j int) bool { return s[i] < s[j] }

func input()(url string){
	fmt.Println("输入小说地址，如：http://www.biqukan.com/1_1094/")
	fmt.Scanln(&url)
	return
}

func main() {
	//url := "http://www.biqukan.com/1_1094/"
	url:=input()
	urls := GetZhang(url)
	//t1:=time.Now()
	urls = []string(Duplicate(urls))
	//sort.Sort(sstrs(urls))
	//fmt.Println(time.Now().Sub(t1))
	os.MkdirAll("download",os.ModeType)
	for _, u := range urls {
		wg.Add(1)
		go DownLoadA("download",Host + u)
		//t,c:=GetZhangContent(Host+u)
		//appendfile("1.txt",t+`\n`+c+`\n`)
		//fmt.Println(u)
	}
	wg.Wait()
}

func DownLoadA(dirname,url string) {
	thch<-true
	t, c := GetZhangContent(url)
	appendfile(fmt.Sprintf("%s\\%s.txt",dirname, t), c+`\n`)
	fmt.Println("下载", t)
	<-thch
	wg.Done()
}

//去除重复
func Duplicate(a []string) (ret []string) {
	m := map[string]bool{}
	for _, v := range a {
		m[v] = false
	}
	ret = make([]string, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	return
}

//追加写入
func appendfile(filename, str string) {
	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(str)
	fd.Write(buf)
	fd.Close()
}

func gentitle(t string)(nt string){
	cnum,err:=mystr.GetBetween(t,"第","章")
	tts:=strings.Split(t," ")
	tt:=strings.Join(tts[len(tts)-1:]," ")
	if err != nil {
		cnum=tts[0]
		nt=t
		return
	}
	nt="第"+cnum+"章"+" "+tt
	return
}

func GetZhangContent(url string) (title, text string) {
	html, _ := nettools.HttpGet(url)
	html = mystr.GBKTOUTF8(html)
	soup := bs.Init(html)
	title = soup.SelByClass("content")[0].SelByTag("h1")[0].Value
	title=gentitle(title)
	text = soup.SelByClass("showtxt")[0].Value
	text = strings.Replace(text, "&nbsp;", "", -1)
	text = strings.Replace(text, "\r\r", "\n", -1)
	return
}

func GetZhang(url string) (urls []string) {
	soup := bs.Init(url)
	nodes := soup.SelByTag("dl")[0].SelByTag("a")
	urls = make([]string, len(nodes))
	for i, n := range nodes {
		urls[i] = (*n.Attrs)["href"]
	}
	return
}
