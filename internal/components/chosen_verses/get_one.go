package chosen_verses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetOneChosenVersesItemInput struct {
	ID string `path:"id" maxLength:"100" doc:"Chosen Verses's ID"`
}

type GetOneChosenVersesItemOutput struct {
	Body   ChosenVerses_Item
	Status int
}

func GetOneChosenVersesHandler(ctx context.Context, input *GetOneChosenVersesItemInput) (*GetOneChosenVersesItemOutput, error) {
	row := database.Pool.QueryRow(ctx, "SELECT id, poet_id, poem_id, verses FROM chosen_verses WHERE id = $1;", input.ID)

	var chosen_verses ChosenVerses_Item
	err := row.Scan(
		&chosen_verses.Id,
		&chosen_verses.Poet_id,
		&chosen_verses.Poem_id,
		&chosen_verses.Verses,
	)

	if err != nil {
		return nil, huma.Error404NotFound("Chosen Verses's not found")
	}

	resp := &GetOneChosenVersesItemOutput{
		Body:   chosen_verses,
		Status: http.StatusOK,
	}

	return resp, nil

}
