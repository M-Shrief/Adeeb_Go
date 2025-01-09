package poets

import (
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetAllPoetsInput struct {
}

type GetAllPoetsOutput struct {
	Body   []sqlc.GetPoetsRow
	Status int
}

func GetAllPoetsHandler(ctx context.Context, input *GetAllPoetsInput) (*GetAllPoetsOutput, error) {
	poets, err := sqlc.Q.GetPoets(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Poets are not available")
	}

	resp := &GetAllPoetsOutput{
		Body:   poets,
		Status: http.StatusOK,
	}

	return resp, nil
}
