package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type student struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Age      int    `json:"age"`
	Location string `json:"location"`
}

var (
	db, _ = sql.Open("mysql", "root:jerrytran97@tcp(127.0.0.1:3306)/goblog")
)

// tạo mới một student
func CreateStudent(c echo.Context) error {
	s := &student{}
	if err := c.Bind(s); err != nil {
		return err
	}
	insertDB, err := db.Prepare("INSERT INTO student(id, fullname, age, location) values (?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insertDB.Exec(s.ID, s.Fullname, s.Age, s.Location)
	return c.JSON(http.StatusCreated, s)
}

// lấy student dựa vào param id
func GetStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	getDB, err := db.Query("SELECT * FROM student WHERE id=", id)
	if err != nil {
		panic(err.Error())
	}
	var s student
	_ = getDB.Scan(&s.ID, &s.Fullname, &s.Age, &s.Location)
	return c.JSON(http.StatusOK, s)
}

// update student
func UpdateStudent(c echo.Context) error {
	s := new(student)
	if err := c.Bind(s); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	updateDB, err := db.Prepare("UPDATE student SET fullname=?, age=?, location=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	updateDB.Exec(s.Fullname, s.Age, s.Location, id)
	return c.JSON(http.StatusOK, s)
}

// xóa một student dựa vào id
func DeleteStudent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	deleteDB, err := db.Prepare("DELETE FROM student WHERE id=?")
	if err != nil {
		panic(err)
	}
	deleteDB.Exec(id)
	return c.NoContent(http.StatusNoContent)
}

// lấy ra tất cả student
func GetStudents(c echo.Context) error {
	var sliceUsers []student
	result, _ := db.Query("SELECT * FROM student")
	for result.Next() {
		var s student
		_ = result.Scan(&s.ID, &s.Fullname, &s.Age, &s.Location)
		sliceUsers = append(sliceUsers, s)
	}
	return c.JSON(http.StatusOK, sliceUsers)
}
