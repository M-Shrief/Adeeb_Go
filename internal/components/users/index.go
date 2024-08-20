package users

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func RegisterAPI(api huma.API) {
	huma.Register(
		api,
		huma.Operation{
			OperationID: "signup-users",
			Method:      http.MethodPost,
			Path:        "/users/signup",
			Summary:     "Signup a User",
			Description: "User signing up.",
			Tags:        []string{"Users"},
		},
		SignupHandler,
	)

	huma.Register(
		api,
		huma.Operation{
			OperationID: "login-users",
			Method:      http.MethodPost,
			Path:        "/users/login",
			Summary:     "User login",
			Description: "User logining in.",
			Tags:        []string{"Users"},
		},
		LoginHandler,
	)
}