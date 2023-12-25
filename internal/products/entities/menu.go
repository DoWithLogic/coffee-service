package entities

import (
	"time"

	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
)

type Menu struct {
	ID               int64      `db:"id"`
	MenuCategoriesID int64      `db:"menu_categories_id"`
	Name             string     `db:"name"`
	Description      string     `db:"description"`
	Price            float64    `db:"price"`
	CreatedAt        time.Time  `db:"created_at"`
	UpdatedAt        *time.Time `db:"updated_at"`
}

type ListMenu []Menu

type UpdateMenu struct {
	ID               int64      `db:"id"`
	MenuCategoriesID int64      `db:"menu_categories_id"`
	Name             string     `db:"name"`
	Description      string     `db:"description"`
	Price            float64    `db:"price"`
	UpdatedAt        *time.Time `db:"updated_at"`
}

func NewMenuEntities(menu *dtos.Menu) *Menu {
	return &Menu{
		ID:               menu.ID,
		MenuCategoriesID: menu.MenuCategoriesID,
		Name:             menu.Name,
		Description:      menu.Description,
		Price:            menu.Price,
		CreatedAt:        time.Now(),
	}
}

func NewMenuDTO(menu *Menu) *dtos.Menu {
	return &dtos.Menu{
		ID:               menu.ID,
		MenuCategoriesID: menu.MenuCategoriesID,
		Name:             menu.Name,
		Description:      menu.Description,
		Price:            menu.Price,
		CreatedAt:        menu.CreatedAt,
	}
}

func NewUpdateMenuEntities(um dtos.UpdateMenu) UpdateMenu {
	var now = time.Now()

	return UpdateMenu{
		ID:               um.ID,
		MenuCategoriesID: um.MenuCategoriesID,
		Name:             um.Name,
		Description:      um.Description,
		Price:            um.Price,
		UpdatedAt:        &now,
	}
}

func NewListMenuDTO(listMenu ListMenu) dtos.ListMenu {
	var listMenuData dtos.ListMenu
	for _, menu := range listMenu {
		listMenuData = append(listMenuData, *NewMenuDTO(&menu))
	}

	return listMenuData
}
