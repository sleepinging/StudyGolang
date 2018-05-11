package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

func Initsql() {
	var err error
	db, err = sql.Open("mysql",
		"root:123456@tcp(127.0.0.1:3306)/mjxt?charset=utf8")
	checkErr(err)
}

func main() {
	Initsql()
	//deltest()
	//addtest()
	//querytest()
	fmt.Println(GetJSON("SELECT * FROM ad"))
	db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func GetJSON(sqlString string) (string, error) {
	stmt, err := db.Prepare(sqlString)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	//fmt.Println(string(jsonData))
	return string(jsonData), nil
}
