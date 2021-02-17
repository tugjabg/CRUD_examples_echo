package main

import (
	"CRUD_examples_echo/service"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	server.POST("/student", service.CreateStudent)
	server.GET("/student/:id", service.GetStudent)
	server.GET("/students", service.GetStudents)
	server.PUT("/student/:id", service.UpdateStudent)
	server.DELETE("/student/:id", service.DeleteStudent)
	//server.POST("/login", genToken, middleware.BasicAuth(basicAuth))
	server.Logger.Fatal(server.Start(":2609"))
}

//
//type user struct {
//	name string `json:"name" form:"name" xml:"name" query:"name"`
//	pass string `json:"pass" form:"pass" xml:"pass" query:"pass"`
//}
//
//type loginResponse struct {
//	Token string `json:"token"`
//}
//
//func (u user) getName() string {
//	return u.name
//}
//
//func (u *user) setName(name string) {
//	u.name = name
//}
//
//func (u user) getPass() string {
//	return u.pass
//}
//
//func (u *user) setPass(pass string) {
//	u.pass = pass
//}
//
//func basicAuth(name string, pass string, c echo.Context) (bool, error) {
//	if (name != "admin") || (pass != "123456") {
//		return false, nil
//	}
//	return true, nil
//}
//
//func genToken(c echo.Context) error {
//	return c.JSON(http.StatusOK, &loginResponse{
//		Token: "123456",
//	})
//}
