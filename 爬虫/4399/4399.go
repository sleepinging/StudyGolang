package main

import (
	"fmt"
	"github.com/pysrc/bs"
	"errors"
	"twt/mystr"
	"strings"
	"path"
	"twt/nettools"
)

func main() {
	url := "http://www.4399.com/flash/193053.htm"
	fmt.Println(DownloadGame(url))
}

func DownloadGame(url string)(err error){
	purl,err:=GetPlayUrl(url)
	if err != nil {
		return
	}
	pt,err:=GetGamePath(purl)
	if err != nil {
		return
	}
	fmt.Println("找到swf文件",pt)
	nameid:=path.Base(url)
	savename:=strings.Split(nameid,".")[0]+".swf"
	fmt.Println("正在下载",savename)
	err=nettools.Download(pt,savename)
	return
}

func GetPlayUrl(gameurl string)(playurl string,err error){
	soup := bs.Init(gameurl)
	if IsH5(soup){
		err=errors.New(`H5游戏不能下载`)
		return
	}
	for _, j := range soup.SelByClass("play") {
		for _,j2:=range j.SelByClass("btn"){
			playurl="http://www.4399.com"+(*j2.Attrs)["href"]
			return
		}
	}
	err=errors.New("未找到")
	return
}

func GetGamePath(playurl string)(swfurl string,err error){
	soup := bs.Init(playurl)
	for _, j := range soup.Sel("", &map[string]string{"language": "javascript"}) {
		path,err:=mystr.GetBetween(j.Value,`_strGamePath="`,`";`)
		if err != nil {
			return path,err
		}
		if strings.HasSuffix(path,".swf"){
			swfurl=`http://sxiao.4399.com/4399swf`+path
		}else{
			path=`http://sxiao.4399.com/4399swf`+path
			swfurl,err=GetSwfUrl(path)
			if err != nil {
				return path,err
			}
		}
		return swfurl,err
	}
	err=errors.New("未找到")
	return
}

func GetSwfUrl(jf string)(swfurl string,err error){
	soup := bs.Init(jf)
	for _, j := range soup.Sel("", &map[string]string{"id": "flashgame"}) {
		for _,node:=range j.Sons{
			value,ok:=(*node.Attrs)["value"]
			if !ok {
				continue
			}
			if strings.HasSuffix(value,`.swf`){
				temp:=strings.Split(jf,"/")
				swfurl=strings.Join(temp[:len(temp)-1],"/")+"/"+value
				return
			}
		}
	}
	err=errors.New("未找到")
	return
}

func IsH5(soup *bs.Soup)(r bool){
	for _, j := range soup.Sel("div", &map[string]string{"class": "spe"}) {
		for _,node:= range j.Sons{
			if strings.Index(node.Value,"H5")!=-1{
				return true
			}
		}
	}
	return
}
