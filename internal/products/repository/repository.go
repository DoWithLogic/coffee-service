package repository

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/entities"
	"github.com/DoWithLogic/coffee-service/pkg/observability"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type repository struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

func NewRepository(db *sqlx.DB) products.Repository {
	return &repository{db: db, log: observability.NewZeroLogHook().Z()}
}

func (r *repository) InsertMenuCategory(ctx context.Context, menuCategory *entities.MenuCategory) error {
	result, err := r.db.NamedExecContext(ctx, insertMenuCategory, menuCategory)
	if err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][InsertMenuCategory]NamedExecContext")

		return err
	}

	menuCategory.ID, err = result.LastInsertId()
	if err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][InsertMenuCategory]LastInsertId")

		return err
	}

	return nil
}

func (r *repository) DetailMenuCategoryByID(ctx context.Context, menuCategoryID int64) (menuCategory entities.MenuCategory, err error) {
	if err := r.db.GetContext(ctx, &menuCategory, detailMenuCategory, menuCategoryID); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][DetailMenuCategoryByID]GetContext")

		return menuCategory, err
	}

	return menuCategory, nil
}

func (r *repository) UpdateMenuCategoryByID(ctx context.Context, menuCategory entities.UpdateMenuCategory) error {
	if _, err := r.db.NamedExecContext(ctx, updateMenuCategory, menuCategory); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][UpdateMenuCategoryByID]NamedExecContext")

		return err
	}

	return nil
}

func (r *repository) MenuCategories(ctx context.Context) (dataMenuCategories entities.MenuCategories, err error) {
	if err := r.db.SelectContext(ctx, &dataMenuCategories, menuCategories); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][MenuCategories]SelectContext")

		return dataMenuCategories, err
	}

	return dataMenuCategories, nil
}

func (r *repository) InsertMenu(ctx context.Context, menu *entities.Menu) error {
	result, err := r.db.NamedExecContext(ctx, insertMenu, menu)
	if err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][InsertMenu]NamedExecContext")

		return err
	}

	menu.ID, err = result.LastInsertId()
	if err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][InsertMenu]LastInsertId")

		return err
	}

	return nil
}

func (r *repository) DetailMenu(ctx context.Context, menuID int64) (menuData entities.Menu, err error) {
	if err := r.db.GetContext(ctx, &menuData, detailMenu, menuID); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][DetailMenu]GetContext")

		return menuData, err
	}

	return menuData, nil
}

func (r *repository) UpdateMenuByID(ctx context.Context, request entities.UpdateMenu) error {
	if _, err := r.db.NamedExecContext(ctx, updateMenu, request); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][UpdateMenuByID]NamedExecContext")

		return err
	}

	return nil
}

func (r *repository) ListMenu(ctx context.Context) (listMenuData entities.ListMenu, err error) {
	if err := r.db.SelectContext(ctx, &listMenuData, listMenu); err != nil {
		r.log.Err(err).Ctx(ctx).Msg("[products][ListMenu]SelectContext")

		return listMenuData, err
	}

	return listMenuData, nil
}
