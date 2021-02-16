package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connect database")
	db, err := sql.Open("mysql", "root:jerrytran97@tcp(127.0.0.1:3306)/demo_schema")
	if err != nil {
		fmt.Println("Kết nối thất bại")
	} else {
		fmt.Println("Kết nối thành công")
	}
	defer db.Close()
	result, err := db.Query("SELECT * FROM student")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for result.Next() {
		var s student
		err := result.Scan(&s.Id, &s.Fullname, &s.Age, &s.Location)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(s)
	}

}

type student struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}
