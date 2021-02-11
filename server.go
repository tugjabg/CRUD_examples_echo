package main

import (
	"github.com/labstack/echo"
	"service"
)

package main

import (
"github.com/labstack/echo"
"service"
)

func main() {
	server := echo.New()
	server.POST("/users", service.CreateUser)
	server.GET("/users/:id", service.GetUser)
	server.PUT("/users/:id", service.UpdateUser)
	server.DELETE("/users/:id", service.DeleteUser)

	server.Logger.Fatal(server.Start(":1323"))
}

//package service
//
//import (
//"github.com/labstack/echo"
//"net/http"
//"strconv"
//)
//
//type user struct {
//	ID int `json:"id"`
//	Name string `json:"name"`
//}
//
//var (
//	users = map[int]*user{}
//	seq   = 1
//)
//
//func CreateUser(c echo.Context) error {
//	u := &user{
//		ID: seq,
//	}
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	users[u.ID] = u
//	seq++
//	return c.JSON(http.StatusCreated, u)
//}
//
//func GetUser(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	return c.JSON(http.StatusOK, users[id])
//}
//
//func UpdateUser(c echo.Context) error {
//	u := new(user)
//	if err := c.Bind(u); err != nil {
//		return err
//	}
//	id, _ := strconv.Atoi(c.Param("id"))
//	users[id].Name = u.Name
//	return c.JSON(http.StatusOK, users[id])
//}
//
//func DeleteUser(c echo.Context) error {
//	id, _ := strconv.Atoi(c.Param("id"))
//	delete(users, id)
//	return c.NoContent(http.StatusNoContent)
//}
