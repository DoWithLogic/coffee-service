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
		r.log.Err(err).Msg("[products][InsertMenuCategory]NamedExecContext")

		return err
	}

	menuCategory.ID, err = result.LastInsertId()
	if err != nil {
		r.log.Err(err).Msg("[products][InsertMenuCategory]LastInsertId")

		return err
	}

	return nil
}

func (r *repository) DetailMenuCategoryByID(ctx context.Context, menuCategoryID int64) (menuCategory entities.MenuCategory, err error) {
	if err := r.db.GetContext(ctx, &menuCategory, detailMenuCategory, menuCategoryID); err != nil {
		r.log.Err(err).Msg("[products][DetailMenuCategoryByID]GetContext")

		return menuCategory, err
	}

	return menuCategory, nil
}

func (r *repository) UpdateMenuCategoryByID(ctx context.Context, menuCategory entities.UpdateMenuCategory) error {
	if _, err := r.db.NamedExecContext(ctx, updateMenuCategory, menuCategory); err != nil {
		r.log.Err(err).Msg("[products][UpdateMenuCategoryByID]NamedExecContext")

		return err
	}

	return nil
}

func (r *repository) MenuCategories(ctx context.Context) (dataMenuCategories entities.MenuCategories, err error) {
	if err := r.db.SelectContext(ctx, &dataMenuCategories, menuCategories); err != nil {
		r.log.Err(err).Msg("[products][MenuCategories]SelectContext")

		return dataMenuCategories, err
	}

	return dataMenuCategories, nil
}
