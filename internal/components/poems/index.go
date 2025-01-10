package poems

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type Verse struct {
	First string `json:"first"`
	Sec   string `json:"sec"`
}

type Poem struct {
	Id      pgtype.UUID `json:"id"`
	Intro   string      `json:"intro"`
	Poet_id pgtype.UUID `json:"poet_id"`
	Verses  []Verse     `json:"verses"`
}

func RegisterAPI(api huma.API) {

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "create-poem",
			Method:        http.MethodPost,
			Path:          "/api/poems",
			Summary:       "Create Poem",
			Description:   "Creating a new Poem",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusCreated,
		},
		CreatePoemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-poems",
			Method:        http.MethodGet,
			Path:          "/api/poems",
			Summary:       "Get Poems",
			Description:   "Get all Poems",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusOK,
		},
		GetAllPoemsHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-poem",
			Method:        http.MethodGet,
			Path:          "/api/poems/{id}",
			Summary:       "Get Poem",
			Description:   "Get a Poem",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusOK,
		},
		GetOnePoemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "update-poem",
			Method:        http.MethodPut,
			Path:          "/api/poems/{id}",
			Summary:       "Update Poem",
			Description:   "Update a Poem",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusAccepted,
		},
		UpdatePoemHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "delete-poem",
			Method:        http.MethodDelete,
			Path:          "/api/poems/{id}",
			Summary:       "Delete Poem",
			Description:   "Delete Poem by ID.",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusAccepted,
		},
		DeletePoemHandler,
	)
}
