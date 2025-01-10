package proses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5"
)

type GetAllProsesInput struct{}

type GetAllProsesOutput struct {
	Body   []Proses_Item
	Status int
}

func GetAllProsesHandler(ctx context.Context, input *GetAllProsesInput) (*GetAllProsesOutput, error) {
	rows, err := database.Pool.Query(ctx, "SELECT id, poet_id, qoute FROM proses;")
	if err != nil {
		return nil, huma.Error404NotFound("Proses are not available", err)
	}
	defer rows.Close()

	proses, err := pgx.CollectRows(rows, pgx.RowToStructByName[Proses_Item])
	if err != nil {
		return nil, huma.Error404NotFound("Proses are not scaned", err)
	}

	return &GetAllProsesOutput{Body: proses, Status: http.StatusOK}, nil
}
