package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreatePoetInput struct {
	Body struct {
		Name       string              `json:"name" minLength:"4" maxLength:"50" example:"عمرو بن كلثوم" doc:"Poet's name"`
		Bio        string              `json:"bio" minLength:"4" maxLength:"500" example:"من أشعر الشعراء بالجاهلية وصاحب المعلقة المشهورة....الخ" doc:"Poet's Bio"`
		TimePeriod database.TimePeriod `json:"time_period" enum:"جاهلي, أموي, عباسي, أندلسي, عثماني ومملوكي, متأخر وحديث" doc:"Poet's Time period"`
	}
}

type CreatePoetOutput struct {
	Body   database.CreatePoetRow
	Status int
}

func CreatePoetHandler(ctx context.Context, input *CreatePoetInput) (*CreatePoetOutput, error) {
	poet, err := database.Q.CreatePoet(
		ctx,
		database.CreatePoetParams{
			Name:       pgtype.Text{String: input.Body.Name, Valid: true},
			Bio:        pgtype.Text{String: input.Body.Bio, Valid: true},
			TimePeriod: database.NullTimePeriod{TimePeriod: input.Body.TimePeriod, Valid: true},
		},
	)

	if err != nil {
		return nil, huma.Error406NotAcceptable("User's data is not acceptable", err) // need to customize errors:[]
	}

	resp := &CreatePoetOutput{
		Body:   poet,
		Status: http.StatusCreated,
	}

	return resp, nil
}
