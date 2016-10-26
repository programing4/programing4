package main

import (
        "database/sql"
        "encoding/json"
	"net/http"
        "fmt"
        "log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
        _ "github.com/go-sql-driver/mysql"
)

type Datalice struct {
        Datas []Data
}               
        
type Data struct {
        Id int
        Name string
        Entry string
        Is_show bool
        Create_at string
}
 
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
        	db := initDB()
                rows, err := db.Query(
                `select * from entries order by created_at;`,
        )   

        var (
                id int 
                name string
                entry string
                is_show bool
                create_at string
        )   

        var a Datalice
        for rows.Next() {
                err := rows.Scan(&id,&name,&entry,&is_show,&create_at)
                if err != nil {
                        log.Fatal(err)
                }
                a.Datas = append(a.Datas, Data{Id: id,Name: name,Entry:entry,Is_show:is_show,Create_at:create_at})
        }

        b, err := json.Marshal(a)
        if err != nil {
                fmt.Println("json err:", err)
        }


	return c.String(http.StatusOK, string(b))
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
