package main

import (
	"twt/nettools"
	"fmt"
	"github.com/axgle/mahonia"
	"twt/mystr"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
func main() {
	url := "http://zjccet.com/plus/chengji.php"
	data := `zkzh=181091121100416&Submit=%B5%C7%C2%BC`
	headers := map[string]string{}
	headers["Referer"] = "http://zjccet.com/plus/chengji.php"
	headers["User-Agent"] = "User-Agent:Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36"
	res, err := nettools.HttpPost(url, data, headers)
	if err != nil {
		fmt.Println("错误", err)
		return
	}
	res = ConvertToString(res, "gbk", "utf-8")
	fmt.Println(res)
	res, err = mystr.GetBetween(res, `成绩：</td>`, `</td>`)
	score, err := mystr.GetBetween(res, `<td>`, ` `)
	fmt.Println(score)
}
