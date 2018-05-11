package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"twt/mystr"
	"twt/nettools"
)

const (
	url    = "http://10.58.16.123/jxxycx.asp"
	dbfile = "info.db"
)

var (
	db   *sql.DB
	stmt *sql.Stmt
)

func gendata(xh int64) (str string) {
	sxh := fmt.Sprintf("%d", xh)
	t := xh / 100000000
	st, et := fmt.Sprintf("%d", t+2), fmt.Sprintf("%d", t+3)
	str = `t_xh=` + sxh + `&t_xm=&s1=` + st + `-` + et + `&s2=1&Submit=%B2%E9%D1%AF`
	return
}

func GetStudentName(xuehao int64) (name string, err error) {
	str, err := nettools.HttpPost(
		url,
		gendata(xuehao),
		nil)
	if err != nil {
		return
	}
	str = mystr.GBKTOUTF8(str)
	name, err = mystr.GetBetween(str, `"textfield2" type="text" class="STYLE4" value="`, `" />`)
	return
}

func init() {
	db, err := sql.Open("sqlite3", dbfile)
	checkErr(err)
	db.Exec("create table if not exists xuehao (xh BIGINT PRIMARY KEY,name VARCHAR (10) NOT NULL);")
	stmt, err = db.Prepare("INSERT INTO xuehao(xh, name) values(?,?)")
	checkErr(err)
}

func InsertInfo(xh int64, name string) (err error) {
	_, err = stmt.Exec(xh, name)
	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
