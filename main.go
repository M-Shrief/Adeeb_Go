package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var err error

type Poet struct {
	ID   string
	Name string
	Bio  string
}

func getPoet(c echo.Context) error {
	id := c.Param("id")
	poet := Poet{}
	stmt, err := db.Prepare("SELECT id, name, bio FROM poet WHERE id = $1")
	row := stmt.QueryRow(id)
	row.Scan(&poet.ID, &poet.Name, &poet.Bio)
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, poet)
}

func getPoets(c echo.Context) error {
	poets := []Poet{}
	err := db.Select(&poets, "SELECT id, name, bio FROM poet")
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, poets)
}

func main() {
	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=adeeb_db_test sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	db.Ping()
	defer db.Close()
	fmt.Println("Database Connected")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/poet/:id", getPoet)
	e.GET("/poets", getPoets)
	e.Logger.Fatal(e.Start(":1323"))
}
