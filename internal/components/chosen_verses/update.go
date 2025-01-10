package chosen_verses

import (
	"Adeeb_Go/internal/components/poems"
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdateChosenVersesInput struct {
	ID   string `path:"id" maxLength:"100" doc:"Chosen Verses's ID"`
	Body struct {
		Poet_ID string        `json:"poet_id" required:"false" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Poem_ID string        `json:"poem_id" required:"false" format:"uuid" maxLength:"100" doc:"Poem's ID"`
		Verses  []poems.Verse `json:"verses" required:"false" maxLength:"1000" doc:"Verses"`
	}
}

type UpdateChosenVersesOutput struct {
	Status int
}

func UpdateChosenVersesHandler(ctx context.Context, input *UpdateChosenVersesInput) (*UpdateChosenVersesOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Invalid ID", err)
	}

	fields := []database.Field{
		{Name: "poet_id", Value: input.Body.Poet_ID},
		{Name: "poem_id", Value: input.Body.Poem_ID},
		{Name: "verses", Value: input.Body.Verses},
	}

	err = database.UpdateQuery("chosen_verses", uuid, fields)
	if err != nil {
		return nil, huma.Error406NotAcceptable("Update not accepted", err)
	}

	return &UpdateChosenVersesOutput{Status: http.StatusAccepted}, nil
}
