package products

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products/entities"
)

type Repository interface {
	InsertMenuCategory(ctx context.Context, menuCategory *entities.MenuCategory) error
	DetailMenuCategoryByID(ctx context.Context, menuCategoryID int64) (menuCategory entities.MenuCategory, err error)
	UpdateMenuCategoryByID(ctx context.Context, menuCategory entities.UpdateMenuCategory) error
	MenuCategories(ctx context.Context) (entities.MenuCategories, error)

	InsertMenu(ctx context.Context, menu *entities.Menu) error
	DetailMenu(ctx context.Context, menuID int64) (entities.Menu, error)
	UpdateMenuByID(ctx context.Context, request entities.UpdateMenu) error
	ListMenu(ctx context.Context) (entities.ListMenu, error)
}
