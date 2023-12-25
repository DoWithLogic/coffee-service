package entities

import (
	"time"

	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
)

type MenuCategory struct {
	ID        int64      `db:"id"`
	Name      string     `db:"name"`
	CreatedAt time.Time  `db:"created_at"`
	UpdateAt  *time.Time `db:"updated_at"`
}

type UpdateMenuCategory struct {
	ID       int64      `db:"id"`
	Name     string     `db:"name"`
	UpdateAt *time.Time `db:"updated_at"`
}

type MenuCategories []MenuCategory

func (m *MenuCategory) IsNull() bool {
	return m.ID == 0
}

func NewMenuCategory(req *dtos.MenuCategory) *MenuCategory {
	return &MenuCategory{
		ID:        req.ID,
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdateAt:  req.UpdateAt,
	}
}

func NewMenuCategoryDTO(req MenuCategory) dtos.MenuCategory {
	return dtos.MenuCategory{
		ID:        req.ID,
		Name:      req.Name,
		CreatedAt: req.CreatedAt,
		UpdateAt:  req.UpdateAt,
	}
}

func NewUpdateMenuCategory(req dtos.UpdateMenuCategoryRequest) UpdateMenuCategory {
	var now = time.Now()

	return UpdateMenuCategory{
		ID:       req.ID,
		Name:     req.Name,
		UpdateAt: &now,
	}
}

func NewMenuCategories(list MenuCategories) dtos.MenuCategories {
	var menuCategories dtos.MenuCategories
	for idx := range list {
		menuCategories = append(menuCategories, NewMenuCategoryDTO(list[idx]))
	}

	return menuCategories
}
