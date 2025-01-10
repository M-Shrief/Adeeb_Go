package proses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetProsesItemInput struct {
	ID string `path:"id" maxLength:"100" doc:"Chosen Verses's ID"`
}

type GetProsesItemOutput struct {
	Body   Proses_Item
	Status int
}

func GetProseItemHandler(ctx context.Context, input *GetProsesItemInput) (*GetProsesItemOutput, error) {
	row := database.Pool.QueryRow(ctx, "SELECT id, poet_id, qoute FROM proses WHERE id = $1;", input.ID)

	var prose_item Proses_Item
	err := row.Scan(
		&prose_item.Id,
		&prose_item.Poet_id,
		&prose_item.Qoute,
	)

	if err != nil {
		return nil, huma.Error404NotFound("Prose item's not found")
	}

	return &GetProsesItemOutput{Body: prose_item, Status: http.StatusOK}, nil
}
