package main

import (
	"net"
	"fmt"
	"regexp"
	"net/url"
	"os"
)

var exp=regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}.*?\[(.*?)\]:(\d+) -> \[(.*?)\]:(\d+) \((\d+|EOF)\):`)

func main() {
	//str:=`Set-Cookie: lid=%E8%AF%A5%E4%BF%A1%E6%81%AF%E6%9A%82%E6%97%B6%E6%9C%AA%E7%9F%A5; Domain=login.taobao.com; Expires=Tue, 20-Aug-2019 09:43:53 GMT; Path=/`
	//r:=tbexp.FindAllSubmatch([]byte(str),1)
	//fmt.Println(url.QueryUnescape(string(r[0][1])))
	file:="/tmp/wacapure.domain"
	_, err := os.Stat(file)
	if err == nil {//文件存在
		err=os.Remove(file)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	netListen, err := net.Listen("unix", file)
	if err != nil {
		fmt.Println(err)
		return
	}
	err=os.Chmod("/tmp/wacapure.domain",777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(l net.Listener) {
		fmt.Println("close")
		l.Close()
	}(netListen)
	fmt.Println("Waiting for clients")
	for {

		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		fmt.Println(conn.RemoteAddr().String(), " request addr conn success")
		go handleConnection(conn)
	}
}

//处理连接
func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "err request addr: ", err)
			return
		}
		//fmt.Println("golang recv:\n", string(buffer[:n]))
		//str:=`2018-08-20 06:46:50 UTC [10.0.0.21]:58438 -> [61.147.223.152]:443 (1598):`
		r:=exp.FindAllSubmatch(buffer[:n],1)
		if len(r)>0&&len(r[0])>5{
			//fmt.Println("src ip:",string(r[0][1]))
			//fmt.Println("src port:",string(r[0][2]))
			//fmt.Println("dst ip:",string(r[0][3]))
			//fmt.Println("dst port:",string(r[0][4]))
			//fmt.Println("len:",string(r[0][5]))
			//fmt.Println("data:",string(buffer[len(r[0][0]):]))
			ana(buffer[len(r[0][0])+1:])
		}/*else{
			//fmt.Println("match failed:",string(buffer[:n]))
		}*/
	}
}

var tbexp=regexp.MustCompile(`(?s)(?i)^GET /avatar/getAvatar.do\?userNick=(.*?)&`)
var tbcookieexp=regexp.MustCompile(`(?s)(?i)set-cookie: lid=(.*?);`)
var wy163exp=regexp.MustCompile(`(?s)(?i)Host: .*163.*Cookie: .*nts_mail_user=(.*?):`)
var mtcookieexp=regexp.MustCompile(`(?s)(?i)Host: .*meituan.*Cookie: unc=(.*?);`)

func ana(payload []byte)(tp uint16,id []byte,ok bool){
	//if bytes.Contains(payload,[]byte("%E8%AF%A5%E4%BF%A1%E6%81%AF%E6%9A%82%E6%97%B6%E6%9C%AA%E7%9F%A5")){
	//	fmt.Println(string(payload))
	//}
	r:=tbexp.FindAllSubmatch(payload,1)
	if len(r)>0{
		u,err:=url.QueryUnescape(string(r[0][1]))
		if err != nil {
			return
		}
		fmt.Println("taobao get:",u)
		return
	}

	r=tbcookieexp.FindAllSubmatch(payload,1)
	if len(r)>0{
		u,err:=url.QueryUnescape(string(r[0][1]))
		if err != nil {
			return
		}
		fmt.Println("taobao cookie:",u)
		return
	}

	r=mtcookieexp.FindAllSubmatch(payload,1)
	if len(r)>0{
		u,err:=url.QueryUnescape(string(r[0][1]))
		if err != nil {
			return
		}
		fmt.Println("meituan cookie:",u)
		return
	}

	r=wy163exp.FindAllSubmatch(payload,1)
	if len(r)>0{
		fmt.Println("wangyi 163::",string(r[0][1]))
		return
	}
	return
}
