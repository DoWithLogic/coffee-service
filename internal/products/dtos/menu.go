package dtos

import (
	"time"

	"github.com/invopop/validation"
)

type Menu struct {
	ID               int64      `json:"id"`
	MenuCategoriesID int64      `json:"menu_categories_id" param:"menu_categories_id" `
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Price            float64    `json:"price"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

type ListMenu []Menu

type UpdateMenu struct {
	ID               int64      `json:"id"`
	MenuCategoriesID int64      `json:"menu_categories_id"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	Price            float64    `json:"price"`
	UpdatedAt        *time.Time `json:"updated_at"`
}

func (um UpdateMenu) HasCategoryID() bool {
	return um.MenuCategoriesID != 0
}

func (x Menu) Validate() error {
	return validation.ValidateStruct(&x,
		validation.Field(&x.MenuCategoriesID, validation.Required),
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Description, validation.Required),
		validation.Field(&x.Price, validation.Required),
	)
}

func (um UpdateMenu) Validate() error {
	return validation.ValidateStruct(&um)
}
