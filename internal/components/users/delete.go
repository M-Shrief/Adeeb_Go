package users

import (
	"Adeeb_Go/internal/auth"
	"Adeeb_Go/internal/database"
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteUserInput struct {
	Auth string `header:"Authorization"`
}

type DeleteUserOutput struct {
	Status int
}

func DeleteUserHandler(ctx context.Context, input *DeleteUserInput) (*DeleteUserOutput, error) {
	claims, err := auth.ValidateToken(
		input.Auth,
		[]string{
			fmt.Sprintf("%v:write", sqlc.RoleManagement),
			fmt.Sprintf("%v:write", sqlc.RoleDBA),
			fmt.Sprintf("%v:write", sqlc.RoleAnalytics),
		},
	)
	if err != nil {
		return nil, huma.Error401Unauthorized("Not Authorizaed")
	}

	userClaims := claims["user"].(map[string]interface{})
	userIDStr := userClaims["id"].(string)
	uuid, err := database.StringToUUID(userIDStr)
	if err != nil {
		return nil, huma.Error404NotFound("User's not found!")
	}

	err = sqlc.Q.DeleteUser(ctx, uuid)
	if err != nil {
		return nil, huma.Error404NotFound("Not deleted", err)
	}

	resp := &DeleteUserOutput{http.StatusAccepted}
	return resp, nil
}
