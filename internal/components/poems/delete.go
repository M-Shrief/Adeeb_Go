package poems

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeletePoemInput struct {
	ID string `path:"id" maxLength:"100" doc:"Poem's ID"`
}

type DeletePoemOutput struct {
	Status int
}

func DeletePoemHandler(ctx context.Context, input *DeletePoemInput) (*DeletePoemOutput, error) {
	// uuid, err := database.StringToUUID(input.ID)
	// if err != nil {
	// 	return nil, huma.Error400BadRequest("Poem's ID is not valid")
	// }

	query := "DELETE FROM poems WHERE id = $1"

	_, err := database.Pool.Exec(ctx, query, input.ID)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Delete not accepted", err)
	}
	return &DeletePoemOutput{Status: http.StatusAccepted}, nil
}
