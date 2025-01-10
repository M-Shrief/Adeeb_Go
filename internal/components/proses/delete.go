package proses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteProsesItemInput struct {
	ID string `path:"id" maxLength:"100" doc:"Prose Item's ID"`
}

type DeleteProsesItemOutput struct {
	Status int
}

func DeleteProseItemHandler(ctx context.Context, input *DeleteProsesItemInput) (*DeleteProsesItemOutput, error) {
	query := "DELETE FROM proses WHERE id = $1"

	_, err := database.Pool.Exec(ctx, query, input.ID)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Delete not accepted", err)
	}
	return &DeleteProsesItemOutput{Status: http.StatusAccepted}, nil

}
