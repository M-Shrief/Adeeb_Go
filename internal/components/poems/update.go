package poems

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdatePoemInput struct {
	ID   string `path:"id" maxLength:"100" doc:"Poem's ID"`
	Body struct {
		Intro   string  `json:"intro" required:"false" minLength:"4" maxLength:"50" example:"حُكمُ المَنِيَّةِ في البَرِيَّةِ جاري" doc:"Poem's intro"`
		Poet_ID string  `json:"poet_id" required:"false" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Verses  []Verse `json:"verses" required:"false" maxLength:"1000" doc:"Poem's Verses"`
	}
}

type UpdatePoemOutput struct {
	Status int
}

func UpdatePoemHandler(ctx context.Context, input *UpdatePoemInput) (*UpdatePoemOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Invalid ID", err)
	}

	fields := []database.Field{
		{Name: "intro", Value: input.Body.Intro},
		{Name: "poet_id", Value: input.Body.Poet_ID},
		{Name: "verses", Value: input.Body.Verses},
	}

	err = database.UpdateQuery("poems", uuid, fields)
	if err != nil {
		return nil, huma.Error406NotAcceptable("Update not accepted", err)
	}

	return &UpdatePoemOutput{Status: http.StatusAccepted}, nil
}
