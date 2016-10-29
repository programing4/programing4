package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

type HttpStatus struct {
	Status error
}

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

	e.GET("/get", handler.get)
	e.POST("/post", handler.post)
	e.PUT("/put", handler.put)
	e.DELETE("/delete", handler.delete)
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

	//insert database
	_, err := handler.db.Query(
		"insert into entries (name,entry) values(?,?);",
		name, entry,
	)

	//create response json
	var stat = HttpStatus{Status: nil}

	if err != nil {
		stat.Status = err
	}

	return c.JSON(http.StatusOK, stat)
}

func (handler *MyHandler) put(c echo.Context) error {
	//request json encoding
	var request_json Data
	body := c.Request().Body()
	decoder := json.NewDecoder(body)
	decoder.Decode(&request_json)

	id := request_json.Id
	name := request_json.Name
	entry := request_json.Entry

	_, err := handler.db.Query(
		"UPDATE entries SET name = ? , entry = ? WHERE id = ?;",
		name, entry, id,
	)

	//create response json
	var stat = HttpStatus{Status: nil}

	if err != nil {
		stat.Status = err
	}

	return c.JSON(http.StatusOK, stat)
}

func (handler *MyHandler) delete(c echo.Context) error {
	var request_json Data
	body := c.Request().Body()
	decoder := json.NewDecoder(body)
	decoder.Decode(&request_json)

	id := request_json.Id

	_, err := handler.db.Query(
		`update entries set is_show=0 where id = ?;`,
		id,
	)
	var stat = HttpStatus{Status: nil}

	if err != nil {
		stat.Status = err
	}
	return c.JSON(http.StatusOK, stat)
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
