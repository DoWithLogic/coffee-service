package products

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
)

type Usecase interface {
	CreateMenuCategory(ctx context.Context, menuCategory *dtos.MenuCategory) error
	DetailMenuCategory(ctx context.Context, menuCategoryID int64) (dtos.MenuCategory, error)
	UpdateMenuCategory(ctx context.Context, request dtos.UpdateMenuCategoryRequest) error
	ListMenuCategory(ctx context.Context) (dtos.MenuCategories, error)

	CreateMenu(ctx context.Context, menu *dtos.Menu) error
	DetailMenu(ctx context.Context, menuID int64) (dtos.Menu, error)
	UpdateMenu(ctx context.Context, request dtos.UpdateMenu) error
	ListMenu(ctx context.Context) (dtos.ListMenu, error)
}
