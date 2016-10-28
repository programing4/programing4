package main

import (
	"database/sql"
	"net/http"
	"os"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type HttpStaus struct {
	status bool
}

type MyError struct {
}

type Datalice struct {
	Datas []Data		`json:"Datas"`
}

type MyHandler struct {
	db *sql.DB
        request_json *Datalice
}

type Data struct {
	Id         int		`json:"id"`
	Name       string	`json:"name"`
	Entry      string	`json:"entry"`
	Is_show    bool  	`json:"is_show"`
	Created_at string	`json:"created_at"`
}

func main() {
	handler := &MyHandler{db: initDB()}
	defer handler.db.Close()
	e := echo.New()

	e.GET("/hello", handler.get)
	e.POST("/world", handler.post)
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
        //request json encoding
        var request_json Data
        body := c.Request().Body()
	decoder := json.NewDecoder(body)
        decoder.Decode(&request_json)

        
        name := request_json.Name
        entry := request_json.Entry

	_, err := handler.db.Query(
		"insert into entries (name,entry) values(?,?);",
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
