package poems

import (
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
)

type GetAllPoemsInput struct {
}

type GetAllPoemsOutput struct {
	Body   []Poem
	Status int
}

func GetAllPoemsHandler(ctx context.Context, input *GetAllPoemsInput) (*GetAllPoemsOutput, error) {
	db := sqlc.GetDBTX()

	rows, err := db.Query(ctx, "SELECT id,intro,poet_id,verses FROM poems;")
	if err != nil {
		return nil, huma.Error404NotFound("Poems are not available", err)
	}
	defer rows.Close()
	poems, err := pgx.CollectRows(rows, pgx.RowToStructByName[Poem])
	if err != nil {
		return nil, huma.Error404NotFound("Poems are not scaned", err)
	}
	return &GetAllPoemsOutput{Body: poems, Status: http.StatusOK}, nil
}
