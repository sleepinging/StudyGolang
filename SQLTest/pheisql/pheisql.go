package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql",
		"root:etonesystem@tcp(27.115.88.42:60040)/wa?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("open sql success!")
	testinsert()
}

func testinsert(){
	sql := "INSERT INTO `test_http` (`deviceId`, `localIp`, `phoneNo`, `domainName`, `name`, `waiip`, `thetime`, `type`, `mac`, `srcPort`, `dstIp`, `dstPort`, `ap_mac`) VALUES ('et203', '192.168', '1588879', 'www.baidu', 'index.html', '125.123', '2018-12-06 09:59:17', '1', '94:88', '34444', '123.123', '80', '78:a0')"
	//stmt, err := db.Prepare(sql)
	_,err:=db.Exec(sql)
	if err != nil {
		fmt.Println(err)
		return
	}
	//_, err = stmt.Exec()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	fmt.Println("insert test success")
	return
}

func main() {
	testinsert()
}
