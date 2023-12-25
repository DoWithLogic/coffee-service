package usecase

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/dtos"
	"github.com/DoWithLogic/coffee-service/internal/products/entities"
	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/rs/zerolog"
)

type usecase struct {
	repo products.Repository
	log  *zerolog.Logger
}

func NewUseCase(r products.Repository, l *zerolog.Logger) products.Usecase {
	return &usecase{repo: r, log: l}
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

func (uc *usecase) CreateMenu(ctx context.Context, menu *dtos.Menu) error {
	menuCategoriesData, err := uc.repo.DetailMenuCategoryByID(ctx, menu.MenuCategoriesID)
	if err != nil {
		return err
	}

	if menuCategoriesData.IsNull() {
		uc.log.Err(apperrors.ErrInvalidMenuCategoriesID).Ctx(ctx).Msg("[products][CreateMenu]menuCategoriesData.IsNull")

		return apperrors.ErrInvalidMenuCategoriesID
	}

	menuEntities := entities.NewMenuEntities(menu)

	err = uc.repo.InsertMenu(ctx, menuEntities)
	if err != nil {
		return err
	}

	menu.ID = menuEntities.ID

	return nil
}

func (uc *usecase) DetailMenu(ctx context.Context, menuID int64) (dtos.Menu, error) {
	var menuData dtos.Menu
	menuEntities, err := uc.repo.DetailMenu(ctx, menuID)
	if err != nil {
		return menuData, err
	}

	menuData = *entities.NewMenuDTO(&menuEntities)

	return menuData, nil
}

func (uc *usecase) UpdateMenu(ctx context.Context, request dtos.UpdateMenu) error {
	if request.HasCategoryID() {
		menuCategoriesData, err := uc.repo.DetailMenuCategoryByID(ctx, request.MenuCategoriesID)
		if err != nil {
			return err
		}

		if menuCategoriesData.IsNull() {
			uc.log.Err(apperrors.ErrInvalidMenuCategoriesID).Ctx(ctx).Msg("[products][UpdateMenu]menuCategoriesData.IsNull")

			return apperrors.ErrInvalidMenuCategoriesID
		}
	}

	if err := uc.repo.UpdateMenuByID(ctx, entities.NewUpdateMenuEntities(request)); err != nil {
		return err
	}

	return nil
}

func (uc *usecase) ListMenu(ctx context.Context) (dtos.ListMenu, error) {
	var listMenu dtos.ListMenu
	listMenuData, err := uc.repo.ListMenu(ctx)
	if err != nil {
		return listMenu, err
	}

	return entities.NewListMenuDTO(listMenuData), nil
}
