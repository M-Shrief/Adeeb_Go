package main

import (
	"encoding/json"
	"fmt"
	"log"
	"my-way/config"
	"my-way/datasource"
	"net/http"
	"time"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Poet struct {
	ID   string `json:"id" redis:"id"` // use redis tags when using redis Hashmaps
	Name string `json:"name" redis:"name"`
	Bio  string `json:"bio" redis:"bio"`
}

func getPoet(c echo.Context) error {
	id := c.Param("id")
	poet := new(Poet)

	rPoetKey := fmt.Sprintf("poet:%v", id)
	val, err := datasource.Redis.Get(c.Request().Context(), rPoetKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &poet)
		fmt.Println("From Redis")
		return c.JSON(http.StatusOK, poet)
	}

	stmt, err := datasource.DB.Prepare("SELECT id, name, bio FROM poet WHERE id = $1")
	row := stmt.QueryRow(id)
	row.Scan(&poet.ID, &poet.Name, &poet.Bio)

	mPoet, _ := json.Marshal(poet)
	ttl := 15 * time.Minute
	datasource.Redis.Set(c.Request().Context(), rPoetKey, mPoet, ttl)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("From Database")
	return c.JSON(http.StatusOK, poet)
}

func getPoets(c echo.Context) error {
	poets := []Poet{}
	err := datasource.DB.Select(&poets, "SELECT id, name, bio FROM poet")
	if err != nil {
		log.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, poets)
}

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variables config.", err)
	}

	config.AppConfig = &conf

	datasource.ConnectDB()
	defer datasource.DB.Close()

	datasource.ConnectRedis()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/poet/:id", getPoet)
	e.GET("/poets", getPoets)
	e.Logger.Fatal(e.Start(":1323"))
}
