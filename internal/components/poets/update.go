package poets

import (
	"Adeeb_Go/internal/database"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type UpdatePoetInput struct {
	ID   string `path:"id" maxLength:"100" doc:"Poet's ID"`
	Body struct {
		Name       string `json:"name" required:"false" minLength:"4" maxLength:"50" example:"عمرو بن كلثوم" doc:"Poet's name"`
		Bio        string `json:"bio" required:"false" minLength:"4" maxLength:"500" example:"من أشعر الشعراء بالجاهلية وصاحب المعلقة المشهورة....الخ" doc:"Poet's Bio"`
		TimePeriod string `json:"time_period,omitempty" required:"false" enum:"جاهلي,أموي,عباسي,أندلسي,عثماني ومملوكي,متأخر وحديث" doc:"Poet's Time period"`
	}
}

type UpdatePoetOutput struct {
	Status int
}

func UpdatePoetHandler(ctx context.Context, input *UpdatePoetInput) (*UpdatePoetOutput, error) {
	uuid, err := database.StringToUUID(input.ID)
	if err != nil {
		return nil, huma.Error422UnprocessableEntity("Invalid ID", err)
	}

	fields := []database.Field{
		{Name: "name", Value: input.Body.Name},
		{Name: "bio", Value: input.Body.Bio},
		{Name: "time_period", Value: input.Body.TimePeriod},
	}

	err = database.UpdateQuery("poets", uuid, fields)

	if err != nil {
		return nil, huma.Error406NotAcceptable("Update not accepted", err)
	}

	return &UpdatePoetOutput{Status: http.StatusAccepted}, nil
}
