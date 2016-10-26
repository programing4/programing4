package main

import (
        "database/sql"
	"net/http"
        "fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
        _ "github.com/go-sql-driver/mysql"
)

type Hello struct {
	Serve int
}

func main() {
	e := echo.New()
        _ = initDB()

	e.GET("/hello", get)
	e.POST("/world", post)
	e.PUT("/put", put)
	e.DELETE("/delete", delete)
	e.Run(standard.New(":4000"))
}

func get(c echo.Context) error {
	text := c.QueryParam("text")
	name := c.QueryParam("name")
	return c.String(http.StatusOK, "Hello, World!: "+text+" "+name)
}

func post(c echo.Context) error {
	text := c.FormValue("text")
	name := c.FormValue("name")
	return c.String(http.StatusOK, "Nice:"+text+" "+name+" !!")
}

func put(c echo.Context) error {
	text := c.FormValue("text")
	return c.String(http.StatusOK, "put:"+text+" !!")
}

func delete(c echo.Context) error {
	is_show := c.QueryParam("is_show")
	return c.String(http.StatusOK, "delete:"+is_show+" !!")
}

func initDB() *sql.DB{
        db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang")

        if err != nil {
            fmt.Print(err)
        }

        return db
}
