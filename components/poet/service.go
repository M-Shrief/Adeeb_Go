package component_poet

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"my-way/datasource"
	"time"
)

func selectPoets(poets *[]Poet) error {
	return datasource.DB.Select(poets, "SELECT id, name, bio FROM poet")
}

func selectPoet(c context.Context, id string, poet *Poet) error {
	rPoetKey := fmt.Sprintf("poet:%v", id)
	fmt.Println(rPoetKey)
	val, err := datasource.Redis.Get(c, rPoetKey).Result()
	if err == nil {
		json.Unmarshal([]byte(val), &poet)
		fmt.Println("From Redis")
		return nil
	}

	stmt, err := datasource.DB.Prepare("SELECT id, name, bio FROM poet WHERE id = $1")
	row := stmt.QueryRow(id)
	row.Scan(&poet.ID, &poet.Name, &poet.Bio)

	mPoet, _ := json.Marshal(&poet)
	ttl := 15 * time.Minute
	datasource.Redis.Set(c, rPoetKey, mPoet, ttl)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("From Database")
	return nil
}
