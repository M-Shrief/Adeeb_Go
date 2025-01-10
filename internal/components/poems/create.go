package poems

import (
	"Adeeb_Go/internal/database/sqlc"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreatePoemInput struct {
	Body struct {
		Intro   string  `json:"intro" minLength:"4" maxLength:"50" example:"حُكمُ المَنِيَّةِ في البَرِيَّةِ جاري" doc:"Poem's intro"`
		Poet_ID string  `json:"poet_id" format:"uuid" maxLength:"100" doc:"Poet's ID"`
		Verses  []Verse `json:"verses" maxLength:"1000" doc:"Poem's Verses"`
	}
}

type CreatePoemOutput struct {
	Body   any
	Status int
}

func CreatePoemHandler(ctx context.Context, input *CreatePoemInput) (*CreatePoemOutput, error) {
	db := sqlc.GetDBTX()
	poem, err := db.Exec(
		ctx,
		"insert into Poems (intro, poet_id, verses) values ($1, $2, $3);",
		input.Body.Intro,
		input.Body.Poet_ID,
		input.Body.Verses,
	)
	if err != nil {
		return nil, huma.Error406NotAcceptable("Poem's data is not acceptable", err) // need to customize errors:[]
	}

	resp := &CreatePoemOutput{
		Body:   poem,
		Status: http.StatusCreated,
	}

	return resp, nil

}
