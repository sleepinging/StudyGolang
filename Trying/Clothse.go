package main

import (
	"twt/nettools"
	"fmt"
	"twt/mystr"
	"encoding/json"
)

type JSONApiSaveOrder struct {
	Success int    `json:"success"`
	Msg     string `json:"msg"`
	Data    Data   `json:"data"`
}

var (
	saves    JSONApiSaveOrder
	tmptoken = ""
)

type Data struct {
	OrderNo         string `json:"orderNo"`
	TimeStamp       string `json:"timeStamp"`
	PaySign         string `json:"paySign"`
	SuccessURL      string `json:"successUrl"`
	OrderID         string `json:"orderId"`
	AType           int    `json:"aType"`
	Money           int    `json:"money"`
	AppointmentType int    `json:"appointmentType"`
	AppID           string `json:"appId"`
	NonceStr        string `json:"nonceStr"`
	Package         string `json:"package"`
	PayMent         string `json:"payMent"`
}

func HomeList() (res string) {
	headers := map[string]string{}
	headers["referer"] = "https://wx.cooleasy.net/?s=84dc720d-16f0-4061-aefa-c955ead48e97&code=061GQnjL1Q2tD41A0QiL1NIEjL1GQnjb&state=1"
	headers["Cookie"] = "ASP.NET_SessionId=clbz2qiwoizujambco11ngh4;" +
		" __RequestVerificationToken=cg7fNPsxsO7wN9LIUQPnkEd99_iM70aypUIoCS2lxH_ipYL9iKEQm7riLmPxG_-Rzmm2S5yYjNla8pNCSY7D68N5u4vcnC3cCg4yF3JA1UY1;" +
		" UM_distinctid=16333b845360-01ec51284e7dfc-2150431a-38400-16333b8453810d;" +
		" CNZZDATA1263640290=1161939214-1525577383-%7C1526127924"
	res, err := nettools.HttpGeth("https://wx.cooleasy.net/Home/PriceList/?machineSn=105806",
		headers)
	if err != nil {
		return
	}
	return
}

func ApiSaveOrder() (res string) {
	headers := map[string]string{}
	headers["referer"] = "https://wx.cooleasy.net/Home/PriceList/?machineSn=105806"
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	headers["Cookie"] = "ASP.NET_SessionId=clbz2qiwoizujambco11ngh4; __RequestVerificationToken=cg7fNPsxsO7wN9LIUQPnkEd99_iM70aypUIoCS2lxH_ipYL9iKEQm7riLmPxG_-Rzmm2S5yYjNla8pNCSY7D68N5u4vcnC3cCg4yF3JA1UY1;" +
		" UM_distinctid=16333b845360-01ec51284e7dfc-2150431a-38400-16333b8453810d;" +
		" CNZZDATA1263640290=1161939214-1525577383-%7C1526123149"
	data := "__RequestVerificationToken=" + tmptoken +
		"&AppointmentType=0" +
		"&PayMent=1" +
		"&TimeDiscountRate=1" +
		"&Money=3" +
		"&PriceId=dadf30bc-0f97-4200-8bf6-573b7d3cccbc" +
		"&MachineSn=105806&Need_A1_CMD=0" +
		"&SleepTime=200&WaitTime=18"
	res, err := nettools.HttpPost("https://wx.cooleasy.net/Home/ApiSaveOrder",
		data,
		headers)
	if err != nil {
		return
	}
	return
}

func PaySuccess() (res string) {
	headers := map[string]string{}
	headers["referer"] = "http://pay.cooleasy.net/CommPay/WeiXinPay/Index?appId=wxe99459883de83d82" +
		"&timeStamp=" + saves.Data.TimeStamp +
		"&nonceStr=" + saves.Data.NonceStr +
		"&package=" + saves.Data.Package +
		"&paySign=" + saves.Data.PaySign +
		"&successUrl=" + saves.Data.SuccessURL
	headers["Cookie"] = "ASP.NET_SessionId=clbz2qiwoizujambco11ngh4; " +
		"__RequestVerificationToken=cg7fNPsxsO7wN9LIUQPnkEd99_iM70aypUIoCS2lxH_ipYL9iKEQm7riLmPxG_-Rzmm2S5yYjNla8pNCSY7D68N5u4vcnC3cCg4yF3JA1UY1; " +
		"UM_distinctid=16333b845360-01ec51284e7dfc-2150431a-38400-16333b8453810d; " +
		"CNZZDATA1263640290=1161939214-1525577383-%7C1526123149"
	res, err := nettools.HttpGeth("https://wx.cooleasy.net/PaySuccess/Index?"+
		"orderId="+ saves.Data.OrderID,
		headers)
	if err != nil {
		return
	}
	return
}

func GetRandomCode() (res string) {
	headers := map[string]string{}
	headers["referer"] = "https://wx.cooleasy.net/PaySuccess/Index?" +
		"orderId=" + saves.Data.OrderID
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	headers["Cookie"] = "ASP.NET_SessionId=clbz2qiwoizujambco11ngh4;" +
		" __RequestVerificationToken=cg7fNPsxsO7wN9LIUQPnkEd99_iM70aypUIoCS2lxH_ipYL9iKEQm7riLmPxG_-Rzmm2S5yYjNla8pNCSY7D68N5u4vcnC3cCg4yF3JA1UY1;" +
		" UM_distinctid=16333b845360-01ec51284e7dfc-2150431a-38400-16333b8453810d;" +
		" CNZZDATA1263640290=1161939214-1525577383-%7C1526123149"
	data := "__RequestVerificationToken=" + tmptoken +
		"&OrderId=" + saves.Data.OrderID +
		"&RunTime=1"
	res, err := nettools.HttpPost("http://wx.cooleasy.net/PaySuccess/GetRandomCode",
		data,
		headers)
	if err != nil {
		return
	}
	return
}

func main() {
	res := HomeList()
	//fmt.Println(res)
	tmptoken, _ = mystr.GetBetween(res, `"__RequestVerificationToken" type="hidden" value="`, `"`)
	fmt.Println("tmptoken:", tmptoken)
	res = ApiSaveOrder()
	json.Unmarshal([]byte(res), &saves)
	//fmt.Println(saves)
	res = PaySuccess()
	tmptoken, _ = mystr.GetBetween(res, `"__RequestVerificationToken" type="hidden" value="`, `"`)
	fmt.Println("tmptoken:", tmptoken)
	res = GetRandomCode()
	fmt.Println(res)
}
