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
	val, rErr := datasource.Redis.Get(c, rPoetKey).Result()
	if rErr == nil {
		json.Unmarshal([]byte(val), &poet)
		return nil
	}

	stmt, dbErr := datasource.DB.Prepare("SELECT id, name, bio FROM poet WHERE id = $1")
	if dbErr != nil {
		log.Println(dbErr)
		return dbErr
	}
	row := stmt.QueryRow(id)
	row.Scan(&poet.ID, &poet.Name, &poet.Bio)

	// if it doesn't exist, it will scan it's field into empty strings
	if poet.ID != id {
		return fmt.Errorf("error: %v", not_Found)
	}
	mPoet, _ := json.Marshal(&poet)
	ttl := 15 * time.Minute
	datasource.Redis.Set(c, rPoetKey, mPoet, ttl)
	return nil
}
