package poems

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetOnePoemInput struct {
	ID string `path:"id" maxLength:"100" doc:"Poem's ID"`
}

type GetOnePoemOutput struct {
	Body   Poem
	Status int
}

func GetOnePoemHandler(ctx context.Context, input *GetOnePoemInput) (*GetOnePoemOutput, error) {
	row := database.Pool.QueryRow(ctx, "SELECT id,intro,poet_id,verses FROM poems WHERE id = $1;", input.ID)

	var poem Poem
	err := row.Scan(&poem.Id, &poem.Intro, &poem.Poet_id, &poem.Verses)
	if err != nil {
		return nil, huma.Error404NotFound("Poem's not found")
	}

	resp := &GetOnePoemOutput{
		Body:   poem,
		Status: http.StatusOK,
	}

	return resp, nil
}
