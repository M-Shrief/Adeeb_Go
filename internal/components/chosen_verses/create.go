package chosen_verses

import (
	"Adeeb_Go/internal/components/poems"
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateChosenVersesInput struct {
	Body struct {
		Poet_ID string        `json:"poet_id" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Poem_ID string        `json:"poem_id" format:"uuid" maxLength:"100" doc:"Poem's ID"`
		Verses  []poems.Verse `json:"verses" maxLength:"1000" doc:"Verses"`
	}
}

type CreateChosenVersesOutput struct {
	Body   any
	Status int
}

func CreateChosenVersesHandler(ctx context.Context, input *CreateChosenVersesInput) (*CreateChosenVersesOutput, error) {
	chosen_verses, err := database.Pool.Exec(
		ctx,
		"INSERT INTO chosen_verses (poet_id, poem_id, verses) values ($1, $2, $3);",
		input.Body.Poet_ID,
		input.Body.Poem_ID,
		input.Body.Verses,
	)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Data is not acceptable", err) // need to customize errors:[]
	}

	return &CreateChosenVersesOutput{Body: chosen_verses, Status: http.StatusCreated}, nil
}
