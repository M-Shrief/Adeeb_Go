package proses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateProsesItemInput struct {
	Body struct {
		Poet_ID string `json:"poet_id" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Qoute   string `json:"qoute" minLength:"4" maxLength:"500" example:"شاعر جاهلي، اشتهر بفصاحته و..." doc:"Prose's qoute"`
	}
}

type CreateProsesItemOutput struct {
	Body   any
	Status int
}

func CreateProsesItemHandler(ctx context.Context, input *CreateProsesItemInput) (*CreateProsesItemOutput, error) {
	prose_item, err := database.Pool.Exec(
		ctx,
		"INSERT INTO proses (poet_id, qoute) VALUES ($1, $2);",
		input.Body.Poet_ID,
		input.Body.Qoute,
	)
	if err != nil {
		return nil, huma.Error406NotAcceptable("Prose item's data is not acceptable", err) // need to customize errors:[]
	}

	resp := &CreateProsesItemOutput{
		Body:   prose_item,
		Status: http.StatusCreated,
	}

	return resp, nil
}
