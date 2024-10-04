package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetOneInput struct {
	ID string `path:"id" maxLength:"100" doc:"Poet's ID"`
}

type GetOneOutput struct {
	Body   database.Poet
	Status int
}

func GetOneHandler(ctx context.Context, input *GetOneInput) (*GetOneOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Poet's ID is not valid")
	}

	poet, err := database.Q.GetPoetById(ctx, uuid)
	if err != nil {
		return nil, huma.Error404NotFound("Poet's not found")
	}

	resp := &GetOneOutput{
		Body:   poet,
		Status: http.StatusOK,
	}

	return resp, nil
}
