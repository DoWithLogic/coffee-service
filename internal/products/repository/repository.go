package repository

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/entities"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) products.Repository {
	return &repository{db: db}
}

func (r *repository) InsertMenuCategory(ctx context.Context, menuCategory *entities.MenuCategory) error {
	result, err := r.db.NamedExecContext(ctx, insertMenuCategory, menuCategory)
	if err != nil {
		return err
	}

	menuCategory.ID, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DetailMenuCategoryByID(ctx context.Context, menuCategoryID int64) (menuCategory entities.MenuCategory, err error) {
	if err := r.db.GetContext(ctx, &menuCategory, detailMenuCategory, menuCategoryID); err != nil {
		return menuCategory, err
	}

	return menuCategory, nil
}

func (r *repository) UpdateMenuCategoryByID(ctx context.Context, menuCategory entities.UpdateMenuCategory) error {
	if _, err := r.db.NamedExecContext(ctx, updateMenuCategory, menuCategory); err != nil {
		return err
	}

	return nil
}

func (r *repository) MenuCategories(ctx context.Context) (dataMenuCategories entities.MenuCategories, err error) {
	if err := r.db.SelectContext(ctx, &dataMenuCategories, menuCategories); err != nil {
		return dataMenuCategories, err
	}

	return dataMenuCategories, nil
}
