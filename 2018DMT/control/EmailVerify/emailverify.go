package EmailVerify

import (
	"../../global"
	"../../tools"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"strings"
	"time"
	"fmt"
)

var (
	Verifytimeout = time.Minute * 10                       //过期时间
	dbname        = global.Config.DbInfo.EmailVerifyDb     //邮箱验证数据库
	dbtype        = global.Config.DbInfo.EmailVerifyDbType //数据库类型
	db            *gorm.DB                                 //数据库连接
)

//邮箱验证码信息结构体
type Verifyinfo struct {
	//邮箱
	Email string `gorm:"primary_key"`

	//验证码
	Code string

	//生成时间
	Gentime time.Time
}

//初始化包
func init() {
	dbname = global.CurrPath + dbname
	//fmt.Println("邮箱验证数据库地址:",dbname)
	tdb, err := gorm.Open(dbtype, dbname)
	tools.PanicErr(err, "邮箱验证数据库初始化")
	db = tdb
	if !db.HasTable(&Verifyinfo{}) {
		db.CreateTable(&Verifyinfo{})
	}
	fmt.Println("邮箱验证数据库初始化完成")
}

//删除过期数据
func DeletePass() {
	db.Where("gentime < ?", time.Now().Add(-Verifytimeout)).Delete(&Verifyinfo{})
}

//使用完毕
func Cleanup() {
	//删除过期数据
	DeletePass()
	//关闭连接
	db.Close()
}

//生成验证码
func GenCode() (code string) {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 6)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result[i] = bytes[r.Intn(len(bytes))]
	}
	return string(result)
}

//更新验证码信息
func UpdateCode(email, code string) {
	info := Verifyinfo{
		Email: email,
	}
	db.Where(&Verifyinfo{Email: email}).First(&info)
	info.Code = code
	info.Gentime = time.Now()
	db.Save(&info)
}

//验证验证码信息
func CheckCode(email, code string) (res bool, errinfo string) {
	code = strings.ToUpper(code)
	errinfo = "验证码过期或错误"
	info := Verifyinfo{
		Email: email,
		Code:  code,
	}
	db.Where(&Verifyinfo{Email: email, Code: code}).First(&info)
	if time.Now().Sub(info.Gentime) < Verifytimeout {
		res, errinfo = true, "验证码正确"
	}
	return
}

//发送验证码到指定邮箱
func SendCode(email string) (code string) {
	code = GenCode()
	err := SendEmail(
		email,
		"感谢您的使用，欢迎加入大学帮",
		genhtml(code),
	)
	tools.ShowErr(err)
	UpdateCode(email, code)
	return
}

//生成验证码邮件的网页
func genhtml(code string) (text string) {
	text = `
		<!DOCTYPE html>
<html>

	<head>
		<meta charset="UTF-8">
		<title></title>
		<meta name="viewport" content="width=device-width,initial-scale=1,minimum-scale=1,maximum-scale=1,user-scalable=no" />
		<link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
		<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
		<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
	</head>

	<body>
		<div class="container">
			<br />
			<div class="jumbotron">
				<h1>欢迎加入 大学帮</h1>
				<p>你的验证码是：</p>
				<p style="width: 100%; margin: 0 auto;">
					<a class="btn btn-primary btn-lg btn-block" href="#" role="button" style="margin: 0 auto;">
					` + code + `
					</a>
				</p>
				<br />
				<p style="float: right;">来自 大学帮团队</p>
			</div>
		</div>
	</body>

</html>
		`
	return
}
