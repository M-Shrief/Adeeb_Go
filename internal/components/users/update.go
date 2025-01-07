package users

import (
	"Adeeb_Go/internal/auth"
	"Adeeb_Go/internal/database"
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateUserInput struct {
	Auth string `header:"Authorization"`
	Body struct {
		Name     string          `json:"name" required:"false" maxLength:"50" example:"John Doe" doc:"User's name"`
		Password string          `json:"password" required:"false" maxLength:"100" example:"P@ssword1" doc:"User's password"`
		Roles    []database.Role `json:"roles" required:"false" enum:"Management,DBA,Analytics" doc:"User's roles"`
	}
}

type UpdateUserOutput struct {
	Status int
}

func UpdateUserHandler(ctx context.Context, input *UpdateUserInput) (*UpdateUserOutput, error) {
	claims, err := auth.ValidateToken(
		input.Auth,
		[]string{
			fmt.Sprintf("%v:write", database.RoleManagement),
			fmt.Sprintf("%v:write", database.RoleDBA),
			fmt.Sprintf("%v:write", database.RoleAnalytics),
		},
	)
	if err != nil {
		return nil, huma.Error401Unauthorized("Not Authorizaed")
	}

	userClaims := claims["user"].(map[string]interface{})
	userIDStr := userClaims["id"].(string)
	uuid, err := database.StringToUUID(userIDStr)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Invalid ID", err)
	}

	var hashedPassword string
	if input.Body.Password != "" {
		hashedPassword, err = auth.Hash(input.Body.Password)
		if err != nil {
			return nil, huma.Error500InternalServerError("Error! Please try again later.")
		}
	}

	err = database.Q.UpdateUser(
		ctx,
		database.UpdateUserParams{
			ID:      uuid,
			Column2: input.Body.Name,
			Column3: hashedPassword,
			Column4: input.Body.Roles,
		},
	)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Update not accepted", err)
	}

	return &UpdateUserOutput{Status: http.StatusAccepted}, nil
}
