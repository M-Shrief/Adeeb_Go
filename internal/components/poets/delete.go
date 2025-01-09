package poets

import (
	"Adeeb_Go/internal/database"
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeletePoetInput struct {
	ID string `path:"id" maxLength:"100" doc:"Poet's ID"`
}

type DeletePoetOutput struct {
	Status int
}

func DeletePoetHandler(ctx context.Context, input *DeletePoetInput) (*DeletePoetOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error400BadRequest("Poet's ID is not valid")
	}

	err = sqlc.Q.DeletePoet(ctx, uuid)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Delete not accepted", err)
	}
	return &DeletePoetOutput{Status: http.StatusAccepted}, nil
}
