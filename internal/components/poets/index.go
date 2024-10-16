package poets

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID:   "create-poet",
			Method:        http.MethodPost,
			Path:          "/api/poets",
			Summary:       "Create Poet",
			Description:   "Creating a new Poet",
			Tags:          []string{"Poets"},
			DefaultStatus: http.StatusCreated,
		},
		CreateHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-poets",
			Method:        http.MethodGet,
			Path:          "/api/poets",
			Summary:       "Get Poets",
			Description:   "Get all Poets",
			Tags:          []string{"Poets"},
			DefaultStatus: http.StatusOK,
		},
		GetAllHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-poet",
			Method:        http.MethodGet,
			Path:          "/api/poets/{id}",
			Summary:       "Get Poet",
			Description:   "Get a Poet",
			Tags:          []string{"Poets"},
			DefaultStatus: http.StatusOK,
		},
		GetOneHandler,
	)
}
