package main

import (
	"github.com/axgle/mahonia"
	"twt/nettools"
	"twt/mystr"
	"strconv"
	"fmt"
)

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func GetMoneyByBank(id, pwd string) (money float64, err error) {
	url := `https://web.zj.icbc.com.cn/easypay/waction.do`
	data := `com.icbc.marmot.core.model.modelname=WapGoodsSure&nosession=0&goodsno=15073&goodsprice=0.01&orderamt=0.01&mtype1=&paytype1=epay&ch=0&goodsct=1&paytype=epay` +
		`&resv1=` + id +
		`&resv2=` + pwd
	res, err := nettools.HttpPost(url, data, nil)
	if err != nil {
		panic(err)
	}
	res = ConvertToString(res, "gbk", "utf-8")
	//fmt.Println(res)
	res, err = mystr.GetBetween(res, `resv4`, `>`)
	res, err = mystr.GetBetween(res, `value="`, `"`)
	//fmt.Println(res)
	money, err = strconv.ParseFloat(res, 64)
	return
}
func main() {
	fmt.Println(GetMoneyByBank(`00125102`, `021338`))
}