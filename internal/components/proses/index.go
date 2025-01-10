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
}