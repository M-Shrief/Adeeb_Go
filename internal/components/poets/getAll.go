package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetAllPoetsInput struct {
}

type GetAllPoetsOutput struct {
	Body   []database.GetPoetsRow
	Status int
}

func GetAllPoetsHandler(ctx context.Context, input *GetAllPoetsInput) (*GetAllPoetsOutput, error) {
	poets, err := database.Q.GetPoets(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Poets are not available")
	}

	resp := &GetAllPoetsOutput{
		Body:   poets,
		Status: http.StatusOK,
	}

	return resp, nil
}
