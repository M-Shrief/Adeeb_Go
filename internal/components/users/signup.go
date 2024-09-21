package users

import (
	"Adeeb_Go/internal/auth"
	"Adeeb_Go/internal/database"
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

type SignupInput struct {
	Body struct {
		Name     string          `json:"name" maxLength:"50" example:"John Doe" doc:"User's name"`
		Password string          `json:"password" maxLength:"100" example:"P@ssword1" doc:"User's password"`
		Roles    []database.Role `json:"roles" enum:"Management,DBA,Analytics" doc:"User's roles"`
	}
}

type SignupOutputBody struct {
	User  UserOutput `json:"user" doc:"User's data"`
	Token string     `json:"token" doc:"JWT token"`
}

type SignupOutput struct {
	Body   SignupOutputBody
	Status int
}

func SignupHandler(ctx context.Context, input *SignupInput) (*SignupOutput, error) {

	hashedPassword, err := auth.Hash(input.Body.Password)
	if err != nil {
		return nil, err
	}

	user, err := database.Q.CreateUser(
		ctx,
		database.CreateUserParams{
			Name:     input.Body.Name,
			Password: hashedPassword,
			Roles:    input.Body.Roles,
		},
	)
	if err != nil {
		return nil, huma.Error406NotAcceptable("User's data is not acceptable", err) // need to customize errors:[]
	}

	token, err := auth.CreateJWT(
		time.Hour,
		auth.JWTUserClaims{
			ID:    database.UUIDToString(user.ID),
			Name:  user.Name,
			Roles: user.Roles,
		},
		auth.NewPermission(user.Roles),
	)
	if err != nil {
		return nil, huma.Error500InternalServerError("Error! please try again later")
	}

	resp := &SignupOutput{
		Body: SignupOutputBody{
			UserOutput{database.UUIDToString(user.ID), user.Name, user.Roles},
			token,
		},
		Status: http.StatusCreated,
	}

	return resp, nil
}
