package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var db *sqlx.DB
var err error
var rdb *redis.Client

type Poet struct {
	ID   string `json:"id" redis:"id"` // use redis tags when using redis Hashmaps
	Name string `json:"name" redis:"name"`
	Bio  string `json:"bio" redis:"bio"`
}

func getPoet(c echo.Context) error {
	id := c.Param("id")
	poet := new(Poet)

	rPoetKey := fmt.Sprintf("poet:%v", id)
	val, err := rdb.Get(c.Request().Context(), rPoetKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &poet)
		fmt.Println("From Redis")
		return c.JSON(http.StatusOK, poet)
	}

	stmt, err := db.Prepare("SELECT id, name, bio FROM poet WHERE id = $1")
	row := stmt.QueryRow(id)
	row.Scan(&poet.ID, &poet.Name, &poet.Bio)

	mPoet, _ := json.Marshal(poet)
	ttl := 15 * time.Minute
	rdb.Set(c.Request().Context(), rPoetKey, mPoet, ttl)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("From Database")
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

	rdb = redis.NewClient(&redis.Options{})
	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Redis Connected")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/poet/:id", getPoet)
	e.GET("/poets", getPoets)
	e.Logger.Fatal(e.Start(":1323"))
}
