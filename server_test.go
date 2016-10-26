package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type Hello struct {
	Serve int
}

func main() {
	e := echo.New()

	e.GET("/hello", get)
	e.POST("/world", post)
	e.PUT("/put", put)
	e.DELETE("/delete", delete)
	e.Run(standard.New(":4000"))
}

func get(c echo.Context) error {
	text := c.QueryParam("text")
	name := c.QueryParam("name")
	return c.JSON(http.StatusOK, "Hello, World!: "+text+" "+name)
}

func post(c echo.Context) error {
	text := c.FormValue("text")
	name := c.FormValue("name")
	return c.JSON(http.StatusOK, "Nice:"+text+" "+name+" !!")
}

func put(c echo.Context) error {
	text := c.FormValue("text")
	return c.JSON(http.StatusOK, "put:"+text+" !!")
}

func delete(c echo.Context) error {
	is_show := c.QueryParam("is_show")
	return c.JSON(http.StatusOK, "delete:"+is_show+" !!")
}
