package usecase

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
	"github.com/DoWithLogic/coffee-service/internal/products/entities"
)

type usecase struct {
	repo products.Repository
}

func NewUseCase(r products.Repository) products.Usecase {
	return &usecase{repo: r}
}

func (uc *usecase) CreateMenuCategory(ctx context.Context, menuCategory *dtos.MenuCategory) error {
	var argsInsert = entities.NewMenuCategory(menuCategory)
	if err := uc.repo.InsertMenuCategory(ctx, argsInsert); err != nil {
		return err
	}

	menuCategory.ID = argsInsert.ID

	return nil
}

func (uc *usecase) DetailMenuCategory(ctx context.Context, menuCategoryID int64) (menuCategory dtos.MenuCategory, err error) {
	menuCategoryData, err := uc.repo.DetailMenuCategoryByID(ctx, menuCategoryID)
	if err != nil {
		return menuCategory, err
	}

	return entities.NewMenuCategoryDTO(menuCategoryData), nil
}

func (uc *usecase) UpdateMenuCategory(ctx context.Context, request dtos.UpdateMenuCategoryRequest) error {
	err := uc.repo.UpdateMenuCategoryByID(ctx, entities.NewUpdateMenuCategory(request))
	if err != nil {
		return err
	}

	return nil
}

func (uc *usecase) ListMenuCategory(ctx context.Context) (dtos.MenuCategories, error) {
	menuCategories, err := uc.repo.MenuCategories(ctx)
	if err != nil {
		return dtos.MenuCategories{}, err
	}

	return entities.NewMenuCategories(menuCategories), nil
}
