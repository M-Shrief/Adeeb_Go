package poets

import (
	"Adeeb_Go/internal/database"
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetOnePoetInput struct {
	ID string `path:"id" maxLength:"100" doc:"Poet's ID"`
}

type GetOnePoetOutput struct {
	Body   sqlc.GetPoetByIdRow
	Status int
}

func GetOnePoetHandler(ctx context.Context, input *GetOnePoetInput) (*GetOnePoetOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Poet's ID is not valid")
	}

	poet, err := sqlc.Q.GetPoetById(ctx, uuid)
	if err != nil {
		return nil, huma.Error404NotFound("Poet's not found")
	}

	resp := &GetOnePoetOutput{
		Body:   poet,
		Status: http.StatusOK,
	}

	return resp, nil
}
