package main

import (
	ypclnt "github.com/yunpian/yunpian-go-sdk/sdk"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
	"strings"
)

var db *gorm.DB
var clnt = ypclnt.New("apikeyxxxx")

type Sms struct {
	Num string
	Send int
}


func init() {
	var err error
	db, err = gorm.Open("sqlite3", "sms.db")
	if err != nil {
		panic("打开数据库失败")
	}
}

func Send(phones []string)(err error){
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = strings.Join(phones,",")
	param[ypclnt.TEXT] = "【云片网】您的验证码是1234"
	//param[ypclnt.TEXT] = "【威尼斯人】您的验证码是1234"
	r:=clnt.Sms().BatchSend(param)
	fmt.Println(r)
	return
}

func GetPhones(limit int)(phones []string,c int,err error){
	smss:=new([]Sms)
	err=db.Model(&Sms{}).Where("send =?",0).Limit(limit).Find(smss).Error
	if err != nil {
		return
	}
	c=len(*smss)
	if c==0{
		return
	}
	phones=make([]string,c)
	for i,p:=range *smss{
		phones[i]=p.Num
	}
	return
}

func startsend(){
	pcs,c,err:=GetPhones(100)
	if err != nil {
		fmt.Println(err)
		return
	}
	for c!=0{
		Send(pcs)
		SetSu(pcs)
		pcs,c,err=GetPhones(100)
	}
}

func SetSu(phones []string)(c int,err error){
	tdb:=db.Begin()
	for _,p :=range phones{
		err=tdb.Model(&Sms{}).Where("num =?",p).Update(map[string]interface{}{"send":1}).Error
		if err != nil {
			tdb.Rollback()
			continue
		}
	}
	tdb.Commit()
	return
}

func test(){
	param := ypclnt.NewParam(2)
	param[ypclnt.MOBILE] = "**"
	param[ypclnt.TEXT] = "【云片网】您的验证码是1234"
	r:=clnt.Sms().SingleSend(param)
	fmt.Println(r)
}

func main() {
	//Send([]string{"",""})
	test()
	//pcs,c,_:=GetPhones(2)
	////SetSu(pcs)
	//fmt.Println(c,pcs)
}
