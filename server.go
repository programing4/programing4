package main

import (
	"database/sql"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type Datalice struct {
	Datas []Data
}

type MyHandler struct {
	db *sql.DB
}

type Data struct {
	Id         int
	Name       string
	Entry      string
	Is_show    bool
	Created_at string
}

func main() {
	handler := &MyHandler{db: initDB()}
	defer handler.db.Close()
	e := echo.New()

	e.GET("/hello", handler.get)
	e.POST("/world/:name/:entry", handler.post)
	e.PUT("/put", put)
	e.DELETE("/delete", delete)
	e.Run(standard.New(":4000"))
}

func (handler *MyHandler) get(c echo.Context) error {
	rows, err := handler.db.Query(
		`select * from entries order by created_at;`,
	)
	if err != nil {
		panic(err)
	}

	var (
		id         int
		name       string
		entry      string
		is_show    bool
		created_at string
	)

	var entries Datalice
	for rows.Next() {
		err := rows.Scan(&id, &name, &entry, &is_show, &created_at)
		if err != nil {
			panic(err)
		}
		entries.Datas = append(entries.Datas, Data{Id: id, Name: name, Entry: entry, Is_show: is_show, Created_at: created_at})
	}

	return c.JSON(http.StatusOK, entries)
}

func (handler *MyHandler) post(c echo.Context) error {
	name := c.Param("name")
	entry := c.Param("entry")

	_, err := handler.db.Query(
		"insert into entries (name,entry) values(?,?)",
		name, entry,
	)

        return err
}

func put(c echo.Context) error {
	text := c.FormValue("text")
	return c.String(http.StatusOK, "put:"+text+" !!")
}

func delete(c echo.Context) error {
	is_show := c.QueryParam("is_show")
	return c.String(http.StatusOK, "delete:"+is_show+" !!")
}

func initDB() *sql.DB {
	user := os.Getenv("MYSQL_USERNAME")
	pass := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", user+":"+pass+"@tcp(localhost:3306)/BBS")

	if err != nil {
		panic(err)
	}

	return db
}

