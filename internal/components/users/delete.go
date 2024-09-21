package users

import (
	"Adeeb_Go/internal/auth"
	"Adeeb_Go/internal/database"
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type DeleteInput struct {
	Auth string `header:"Authorization"`
}

type DeleteOutput struct {
	Status int
}

func DeleteHandler(ctx context.Context, input *DeleteInput) (*DeleteOutput, error) {
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
		return nil, huma.Error404NotFound("User's not found!")
	}

	err = database.Q.DeleteUser(ctx, uuid)
	if err != nil {
		return nil, huma.Error404NotFound("Not deleted", err)
	}

	resp := &DeleteOutput{http.StatusAccepted}
	return resp, nil
}
