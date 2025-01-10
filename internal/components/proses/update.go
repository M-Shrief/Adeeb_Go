package proses

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateProsesItemInput struct {
	ID   string `path:"id" maxLength:"100" doc:"Prose Item's ID"`
	Body struct {
		Poet_ID string `json:"poet_id" required:"false" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Qoute   string `json:"qoute" required:"false" minLength:"4" maxLength:"500" example:"شاعر جاهلي، اشتهر بفصاحته و..." doc:"Prose's qoute"`
	}
}

type UpdateProsesItemOutput struct {
	Status int
}

func UpdateProseItemHandler(ctx context.Context, input *UpdateProsesItemInput) (*UpdateProsesItemOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Invalid ID", err)
	}

	fields := []database.Field{
		{Name: "poet_id", Value: input.Body.Poet_ID},
		{Name: "qoute", Value: input.Body.Qoute},
	}

	err = database.UpdateQuery("proses", uuid, fields)
	if err != nil {
		return nil, huma.Error406NotAcceptable("Update not accepted", err)
	}

	return &UpdateProsesItemOutput{Status: http.StatusAccepted}, nil
}
