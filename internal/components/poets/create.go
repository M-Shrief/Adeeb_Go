package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type CreateInput struct {
	Body struct {
		Name       string              `json:"name" minLength:"4" maxLength:"50" example:"عمرو بن كلثوم" doc:"Poet's name"`
		Bio        string              `json:"bio" minLength:"4" maxLength:"500" example:"من أشعر الشعراء بالجاهلية وصاحب المعلقة المشهورة....الخ" doc:"Poet's Bio"`
		TimePeriod database.TimePeriod `json:"time_period" enum:"جاهلي, أموي, عباسي, أندلسي, عثماني ومملوكي, متأخر وحديث" doc:"Poet's Time period"`
	}
}

type CreateOutput struct {
	Body   database.CreatePoetRow
	Status int
}

func CreateHandler(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	poet, err := database.Q.CreatePoet(
		ctx,
		database.CreatePoetParams{
			Name:       input.Body.Name,
			Bio:        input.Body.Bio,
			TimePeriod: input.Body.TimePeriod,
		},
	)

	if err != nil {
		return nil, huma.Error406NotAcceptable("User's data is not acceptable", err) // need to customize errors:[]
	}

	resp := &CreateOutput{
		Body:   poet,
		Status: http.StatusCreated,
	}

	return resp, nil
}
