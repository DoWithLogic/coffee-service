package dtos

import (
	"time"

	"github.com/invopop/validation"
	_ "github.com/invopop/validation"
)

type MenuCategory struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
}

type MenuCategories []MenuCategory

type UpdateMenuCategoryRequest struct {
	ID   int64  `param:"id"`
	Name string `json:"name"`
}

func (x MenuCategory) Validate() error {
	return validation.ValidateStruct(&x,
		validation.Field(&x.Name, validation.Required),
	)
}

func (x UpdateMenuCategoryRequest) Validate() error {
	return validation.ValidateStruct(&x,
		validation.Field(&x.ID, validation.Required),
		validation.Field(&x.Name, validation.Required),
	)
}
