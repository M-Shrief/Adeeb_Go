package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type GetAllInput struct {
}

type GetAllOutput struct {
	Body   []database.GetPoetsRow
	Status int
}

func GetAllHandler(ctx context.Context, input *GetAllInput) (*GetAllOutput, error) {
	poets, err := database.Q.GetPoets(ctx)
	if err != nil {
		return nil, huma.Error404NotFound("Poets are not available")
	}

	resp := &GetAllOutput{
		Body:   poets,
		Status: http.StatusOK,
	}

	return resp, nil
}
