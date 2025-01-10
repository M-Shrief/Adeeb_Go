package chosen_verses

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(api huma.API) {

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "create-chosen_verses-item",
			Method:        http.MethodPost,
			Path:          "/api/chosen_verses",
			Summary:       "Create a Chosen Verses Item",
			Description:   "Creating a new Chosen Verses's item",
			Tags:          []string{"Chosen_Verses"},
			DefaultStatus: http.StatusCreated,
		},
		CreateChosenVersesHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-chosen_verses",
			Method:        http.MethodGet,
			Path:          "/api/chosen_verses",
			Summary:       "Get All Chosen Verses",
			Description:   "Get all Chosen Verses",
			Tags:          []string{"Chosen_Verses"},
			DefaultStatus: http.StatusOK,
		},
		GetAllChosenVersesHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "get-chosen_verses-item",
			Method:        http.MethodGet,
			Path:          "/api/chosen_verses/{id}",
			Summary:       "Get a Chosen Verses Item",
			Description:   "Get a Chosen Verses Item By ID.",
			Tags:          []string{"Chosen_Verses"},
			DefaultStatus: http.StatusOK,
		},
		GetOneChosenVersesHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "update-chosen_verses-item",
			Method:        http.MethodPut,
			Path:          "/api/chosen_verses/{id}",
			Summary:       "Update a Chosen Verses Item",
			Description:   "Update a Chosen Verses Item By ID.",
			Tags:          []string{"Chosen_Verses"},
			DefaultStatus: http.StatusAccepted,
		},
		UpdateChosenVersesHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID:   "delete-chosen_verses-item",
			Method:        http.MethodDelete,
			Path:          "/api/chosen_verses/{id}",
			Summary:       "Delete a Chosen Verses Item",
			Description:   "Delete a Chosen Verses Item by ID.",
			Tags:          []string{"Poems"},
			DefaultStatus: http.StatusAccepted,
		},
		DeleteChosenVersesHandler,
	)
}
