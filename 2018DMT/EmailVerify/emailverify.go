package EmailVerify

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"math/rand"
	"time"
	"twt/mytools"
)

var (
	Verifytimeout = time.Minute * 10 //过期时间
	//infos         []Verifyinfo       //所有验证码信息
	dbname = `data\emailverify.db` //邮箱验证数据库
	db     *gorm.DB                //数据库连接
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

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

//初始化包
func init() {
	pt, _ := mytools.GetCurrentPath()
	dbname = pt + dbname
	tdb, err := gorm.Open("sqlite3", dbname)
	checkerr(err)
	db = tdb
	if !db.HasTable(&Verifyinfo{}) {
		db.CreateTable(&Verifyinfo{})
	}
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
func SendCode(email string) {
	code := GenCode()
	text := `验证码是<br>` +
		code + `<br>` +
		`有效期10分钟`
	SendEmail(email, "验证码", text)
	UpdateCode(email, code)
}
