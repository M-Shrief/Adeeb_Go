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
		CreatePoetHandler,
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
		GetAllPoetsHandler,
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
		GetOnePoetHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "update-poet",
			Method:        http.MethodPut,
			Path:          "/api/poets/{id}",
			Summary:       "Update Poet",
			Description:   "Update a Poet",
			Tags:          []string{"Poets"},
			DefaultStatus: http.StatusAccepted,
		},
		UpdatePoetHandler,
	)
}
