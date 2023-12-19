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
}
