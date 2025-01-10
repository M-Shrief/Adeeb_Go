package chosen_verses

import (
	"Adeeb_Go/internal/components/poems"
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type ChosenVerses struct {
	Id      pgtype.UUID   `json:"id"`
	Poet_id pgtype.UUID   `json:"poet_id"`
	Poem_id pgtype.UUID   `json:"poem_id"`
	Verses  []poems.Verse `json:"verses"`
}

type GetAllChosenVersesInput struct{}

type GetAllChosenVersesOutput struct {
	Body   any
	Status int
}

func GetAllChosenVersesHandler(ctx context.Context, input *GetAllChosenVersesInput) (*GetAllChosenVersesOutput, error) {
	rows, err := database.Pool.Query(ctx, "SELECT id, poet_id, poem_id, verses FROM chosen_verses;")
	if err != nil {
		return nil, huma.Error404NotFound("Poems are not available", err)
	}
	defer rows.Close()

	chosen_verses, err := pgx.CollectRows(rows, pgx.RowToStructByName[ChosenVerses])
	if err != nil {
		return nil, huma.Error404NotFound("Poems are not scaned", err)
	}
	return &GetAllChosenVersesOutput{Body: chosen_verses, Status: http.StatusOK}, nil
}
