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

}
