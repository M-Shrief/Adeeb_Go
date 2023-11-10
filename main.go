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
}

func getPoet(c echo.Context) error {
	// Poet ID from path `Poets/:id`
	id := c.Param("id")
	poets := []Poet{}
	db.Select(&poets, "SELECT id, name FROM poet WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, poets)
}

func getPoets(c echo.Context) error {
	poets := []Poet{}
	db.Select(&poets, "SELECT id, name FROM poet")
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, poets)
}

func main() {
	db, err = sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=adeeb_db_test sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	fmt.Println("Database Connected")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/poet/:id", getPoet)
	// e.GET("/poets", func(ctx echo.Context) error {
	// 	poets := []Poet{}
	// 	db.Select(&poets, "SELECT id, name FROM poet")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return ctx.JSON(http.StatusOK, poets)
	// })
	e.GET("/poets", getPoets)
	e.Logger.Fatal(e.Start(":1323"))
}
