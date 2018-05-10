package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	db *sql.DB
)

func querytest() {
	rows, err := db.Query("SELECT * FROM ad")
	checkErr(err)

	for rows.Next() {
		var id int
		var name string
		var pwd string
		var phone string
		err = rows.Scan(&id, &name, &pwd, &phone)
		checkErr(err)
		fmt.Print(id, "\t")
		fmt.Print(name, "\t")
		fmt.Print(pwd, "\t")
		fmt.Println(phone, "\t")
	}
}

func deltest() {
	sql := "DELETE FROM `ad` WHERE (`id`=?)"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec(2)
	checkErr(err)
}

func addtest() {
	sql := "INSERT INTO `ad` (`id`, `name`, `password`, `phone`) VALUES ('2', '123', '123', '123')"
	stmt, err := db.Prepare(sql)
	checkErr(err)
	_, err = stmt.Exec()
	checkErr(err)
}

func main() {
	var err error
	db, err = sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/mjxt?charset=utf8")
	defer db.Close()
	checkErr(err)
	//deltest()
	addtest()
	querytest()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
