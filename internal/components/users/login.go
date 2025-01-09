package users

import (
	"Adeeb_Go/internal/auth"
	"Adeeb_Go/internal/database"
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

type LoginInput struct {
	Body struct {
		Name     string `json:"name" maxLength:"50" example:"John Doe" doc:"User's name"`
		Password string `json:"password" maxLength:"100" example:"P@ssword1" doc:"User's password"`
	}
}

type LoginOutputBody struct {
	User  UserOutput `json:"user" doc:"User's data"`
	Token string     `json:"token" doc:"JWT token"`
}

type LoginOutput struct {
	Body   LoginOutputBody
	Status int
}

func LoginHandler(ctx context.Context, input *LoginInput) (*LoginOutput, error) {
	user, err := sqlc.Q.GetUserByName(
		ctx,
		input.Body.Name,
	)
	if err != nil {
		return nil, huma.Error404NotFound("User is not found", err) // need to customize errors:[]
	}

	err = auth.CompareHash(user.Password, input.Body.Password)
	if err != nil {
		return nil, huma.Error401Unauthorized("Password is incorrect", err) // need to customize errors:[]
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

	resp := &LoginOutput{
		Body: LoginOutputBody{
			UserOutput{database.UUIDToString(user.ID), user.Name, user.Roles},
			token,
		},
		Status: http.StatusAccepted,
	}

	return resp, nil
}
