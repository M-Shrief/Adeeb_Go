package proses

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type Proses_Item struct {
	Id      pgtype.UUID `json:"id"`
	Poet_id pgtype.UUID `json:"poet_id"`
	Qoute   string      `json:"qoute"`
}

func RegisterAPI(api huma.API) {

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "create-proses-item",
			Method:        http.MethodPost,
			Path:          "/api/proses",
			Summary:       "Create a Proses Item",
			Description:   "Creating a new Proses's item",
			Tags:          []string{"Proses"},
			DefaultStatus: http.StatusCreated,
		},
		CreateProsesItemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-proses",
			Method:        http.MethodGet,
			Path:          "/api/proses",
			Summary:       "Get All Proses",
			Description:   "Get all Proses",
			Tags:          []string{"Proses"},
			DefaultStatus: http.StatusOK,
		},
		GetAllProsesHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-proses-item",
			Method:        http.MethodGet,
			Path:          "/api/proses/{id}",
			Summary:       "Get a Proses Item",
			Description:   "Get a Proses Item By ID.",
			Tags:          []string{"Proses"},
			DefaultStatus: http.StatusOK,
		},
		GetProseItemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "update-proses-item",
			Method:        http.MethodPut,
			Path:          "/api/proses/{id}",
			Summary:       "Update a Proses Item",
			Description:   "Update a Proses Item By ID.",
			Tags:          []string{"Proses"},
			DefaultStatus: http.StatusAccepted,
		},
		UpdateProseItemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "delete-proses-item",
			Method:        http.MethodDelete,
			Path:          "/api/proses/{id}",
			Summary:       "Delete a Proses Item",
			Description:   "Delete a Proses Item by ID.",
			Tags:          []string{"Proses"},
			DefaultStatus: http.StatusAccepted,
		},
		DeleteProseItemHandler,
	)
}
