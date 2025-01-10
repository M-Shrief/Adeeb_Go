package chosen_verses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteChosenVersesInput struct {
	ID string `path:"id" maxLength:"100" doc:"Chosen Verses's ID"`
}

type DeleteChosenVersesOutput struct {
	Status int
}

func DeleteChosenVersesHandler(ctx context.Context, input *DeleteChosenVersesInput) (*DeleteChosenVersesOutput, error) {
	query := "DELETE FROM chosen_verses WHERE id = $1"

	_, err := database.Pool.Exec(ctx, query, input.ID)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Delete not accepted", err)
	}
	return &DeleteChosenVersesOutput{Status: http.StatusAccepted}, nil

}
