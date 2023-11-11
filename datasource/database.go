package datasource

import (
	"fmt"
	"log"
	"my-way/config"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB
var err error

func ConnectDB() {
	sqlxStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v sslmode=disable", config.AppConfig.DB_HOST, config.AppConfig.DB_PORT, config.AppConfig.DB_USER, config.AppConfig.DB_NAME)
	DB, err = sqlx.Connect("postgres", sqlxStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB.Ping()
	fmt.Println("Database Connected")
}
