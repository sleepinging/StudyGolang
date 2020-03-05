package main

import (
	"errors"
	"fmt"
	"net/url"
	"os/exec"
	"regexp"
	"strings"
	"twt/mytools"
	"twt/nettools"
	"sync"
	"os"
	"flag"
)

func t() {
	str := `a=1,a=2,a=3`
	exp := regexp.MustCompile(`a=(\d+)`)
	fmt.Println(exp.FindAllStringSubmatch(str, -1))

	str = `<em>睡觉</em>叫醒系统`
	exp = regexp.MustCompile(`<.*?>`)
	b := exp.ReplaceAll([]byte(str), []byte(""))
	fmt.Println(string(b))
}

type ZhuanLi struct {
	Url         string //简介链接
	DownLoadUrl string //下载链接
	Title       string //标题
	Type        string //类型
	Number      string //专利号
	try         int    //尝试解析下载链接次数
}

//获取搜索结果中的专利链接，标题，类型，专利号
var divzlexp = regexp.MustCompile(`(?s)(?i)<div class="ResultCont">.*?<div class="title">.*?<strong>\[专利\]</strong>.*?<a href="(.*?)".*?>(.*?)</a>.*?<div class="ResultMoreinfo">.*?<span.*?>(.*?)</span>.*?<span.*?>(.*?)</span>`)

////获取搜索结果页数
//var pagesexp=regexp.MustCompile(``)

//获取下载链接
var dlexp = regexp.MustCompile(`<a onclick="upload\('(.*?)','(.*?)','(.*?)','(.*?)','(.*?)','(.*?)','(.*?)'\)"`)

//删除标题中的em标签
var ttexp = regexp.MustCompile(`<.*?>`)

//完整专利信息的通道
var zlch = make(chan ZhuanLi, 100)

//没有下载地址的专利信息
var izlch = make(chan ZhuanLi, 100)

//写入文件信息的通道
var saveinfoch = make(chan string,100)

var wg *sync.WaitGroup

func gettttext(htitle string) (title string) {
	b := ttexp.ReplaceAll([]byte(htitle), []byte(""))
	return string(b)
}

func getdivzls(name string,num int) (zls []ZhuanLi, err error) {
	url := `http://www.wanfangdata.com.cn/search/searchList.do?searchType=patent&pageSize=`+
		fmt.Sprintf("%d",num)+
		`&page=1&searchWord=` +
		url.QueryEscape(name) +
		`&order=correlation&showType=detail&isCheck=check&isHit=&isHitUnit=&firstAuthor=false&rangeParame=all`
	res, err := nettools.HttpGet(url)
	if err != nil {
		//fmt.Println(err)
		return
	}
	//fmt.Println(res)
	matches := divzlexp.FindAllStringSubmatch(res, -1)
	zls = make([]ZhuanLi, len(matches))
	for i, vs := range matches {
		zls[i].Url = `http://www.wanfangdata.com.cn` + vs[1]
		zls[i].Title = gettttext(vs[2])
		zls[i].Type = vs[3]
		zls[i].Number = vs[4]
		fmt.Println("找到专利", zls[i].Title, zls[i].Number, zls[i].Type)
		//fmt.Println(`http://www.wanfangdata.com.cn`+vs[1], gettttext(vs[2]), vs[3], vs[4])
	}
	return
}

//获取js中upload的返回值
func getjsupload(page_cnt, id, language, source_db, title, isoa, tp string) string {
	title = url.QueryEscape(title)
	return "http://www.wanfangdata.com.cn/search/downLoad.do?page_cnt=" + page_cnt + "&language=" + language + "&resourceType=" + tp + "&source=" + source_db + "&resourceId=" + id + "&resourceTitle=" + title + "&isoa=" + isoa + "&type=" + tp
}

//去除末尾指定字符串
func myTrimSuffix(str string, ts []string) (s string) {
	s = str
LABEL_TRIM:
	for {
		for _, v := range ts {
			if strings.HasSuffix(s, v) {
				s = strings.TrimSuffix(s, v)
				continue LABEL_TRIM
			}
		}
		break
	}
	return
}

func getdownloadurl(zlurl string) (url string, err error) {
	res, err := nettools.HttpGet(zlurl)
	if err != nil {
		return
	}
	matches := dlexp.FindAllStringSubmatch(res, -1)
	if len(matches) <= 0 {
		err = errors.New("未找到下载链接")
		return
	}
	furl := getjsupload(matches[0][1], matches[0][2], matches[0][3], matches[0][4], matches[0][5], matches[0][6], matches[0][7])
	//fmt.Println("debug:",furl)
	p, _ := mytools.GetCurrentPath()
	//cmd:=exec.Command(p+"run.bat","getdlurl.js",`"`+url+`"`)
	cmd := exec.Command(p+"getdlurl.exe", p+"phantomjs.exe", p+"getdlurl.js", furl)
	//err:=cmd.Run()
	//cmd:=exec.Command("ping","www.baidu.com")
	buf, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	url = string(buf)
	url = myTrimSuffix(url, []string{" ", "\r", "\n"})
	return
}

//从通道中取出专利下载
func DownLoadTh() {
	wg.Add(1)
	p, _ := mytools.GetCurrentPath()
	for {
		z := <-zlch
		savep := p + z.Title + ".pdf"
		fmt.Println("正在下载", z.Title, "=>", savep)
		nettools.Download(z.DownLoadUrl, savep)
		fmt.Println("下载",z.Title,"成功")
	}
	wg.Done()
}

//更新不完整专利信息放入完整信息的通道
func UpDateTh() {
	wg.Add(1)
	var err error
	for {
		z := <-izlch
		fmt.Println("正在解析", z.Title, "的下载链接...")
		z.DownLoadUrl, err = getdownloadurl(z.Url)
		if err != nil { //出错
			fmt.Println("解析", z.Title, "的下载链接失败",err)
			if z.try<3{	//尝试3次
				z.try++
				izlch <- z//放回去，等待下次解析
			}

		} else { //正常，放入完整通道
			fmt.Println("解析", z.Title, "的下载链接成功")
			zlch <- z
		}
	}
	wg.Done()
}

func appendfile(filename,str_content string)  {
	fd,_:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	buf:=[]byte(str_content)
	fd.Write(buf)
	fd.Close()
}

//保存到文件信息
func saveinfoTh(){
	wg.Add(1)
	for;;{
		t:=<-saveinfoch
		appendfile("infos.txt","\n"+t)
	}
	wg.Done()
}

//往不完整专利信息通道放入数据,还有保存的信息
func GetZLs(zls []ZhuanLi) {
	for _, v := range zls {
		izlch <- v
		t:=fmt.Sprintf("%-020s%-020s%-020s",v.Title,v.Number,v.Type)
		saveinfoch<-t
	}
}

func init() {
	wg=new(sync.WaitGroup)
	t:=fmt.Sprintf("%-020s%-020s%-020s","专利名称","专利号","类型")
	saveinfoch<-t
}

func main() {
	name := flag.String("name", "", "要搜索的专利名字")
	num := flag.Int("num", 10, "要下载的专利数量,目前最大支持50")
	flag.Parse()

	if *name==""{
		fmt.Println("缺少参数name,搜索的专利名字,比如-name=\"睡觉\" -num=20")
		return
	}

	go UpDateTh()
	go DownLoadTh()
	go saveinfoTh()

	zls, err := getdivzls(*name,*num)
	if err != nil {
		panic(err)
	}
	GetZLs(zls)
	//等待
	wg.Wait()
	//url := `http://www.wanfangdata.com.cn/search/downLoad.do?page_cnt=13&language=&resourceType=patent&source=WF&resourceId=CN201710644621.7&resourceTitle=%E7%9D%A1%E8%A7%89%E5%8F%AB%E9%86%92%E7%B3%BB%E7%BB%9F&isoa=&type=patent`
	//url, _ = getdownloadurl(`http://www.wanfangdata.com.cn/details/detail.do?_type=patent&id=CN201710644621.7`)
	//fmt.Println(url)
	//nettools.Download(url, "E:\\1.pdf")
}
