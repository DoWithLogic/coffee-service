package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/products"
	"github.com/DoWithLogic/coffee-service/internal/products/entities"
	"github.com/DoWithLogic/coffee-service/internal/products/repository"
	"github.com/DoWithLogic/coffee-service/pkg/databases"
	"github.com/go-faker/faker/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	cfg config.Config
	db  *sqlx.DB
)

const (
	MIGRATION_PATH = "file://../../../database/migrations"
)

func init() {
	var err error

	cfg, err = config.LoadConfigPath("../../../config/config.integration.test")
	if err != nil {
		panic(errors.Wrap(err, "config.LoadConfigPath"))
	}

	db, err = databases.NewMySQLDB(context.Background(), cfg.Database)
	if err != nil {
		panic(errors.Wrap(err, cfg.Database.DBName))
	}

	sqlDB := db.DB

	err = migrationUP(MIGRATION_PATH, sqlDB)
	if err != nil {
		panic(err)
	}
}

func migrationUP(path string, db *sql.DB) error {
	// Set the database instance for the "mysql" driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	// Set up the migrate instance with the MySQL database driver and file source
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		return err
	}

	// Apply migrations
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func migrationDOWN(path string, db *sql.DB) error {
	// Set the database instance for the "mysql" driver
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	// Set up the migrate instance with the MySQL database driver and file source
	m, err := migrate.NewWithDatabaseInstance(path, "mysql", driver)
	if err != nil {
		return err
	}

	defer m.Close()

	// Apply migrations
	err = m.Down()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

type repositoryTestSuite struct {
	repo products.Repository
	suite.Suite
}

func TestSuiteRepository(t *testing.T) {
	defer func() {
		// err := migrationDOWN(MIGRATION_PATH, db.DB)
		// require.NoError(t, err)
	}()

	suite.Run(t, &repositoryTestSuite{repo: repository.NewRepository(db)})
}

func (r *repositoryTestSuite) Test_products_repositories() {
	var (
		ctx = context.Background()
	)

	argsInsert := &entities.MenuCategory{
		Name:      faker.Name(),
		CreatedAt: time.Now(),
	}

	err := r.repo.InsertMenuCategory(ctx, argsInsert)
	r.Assert().NoError(err)

	detail, err := r.repo.DetailMenuCategoryByID(ctx, argsInsert.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(argsInsert.Name, detail.Name)

	var now = time.Now()

	argsUpdate := entities.UpdateMenuCategory{
		ID:       argsInsert.ID,
		Name:     faker.Name(),
		UpdateAt: &now,
	}

	err = r.repo.UpdateMenuCategoryByID(ctx, argsUpdate)
	r.Assert().NoError(err)

	detailAfterUpdate, err := r.repo.DetailMenuCategoryByID(ctx, argsInsert.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(argsUpdate.Name, detailAfterUpdate.Name)

	menuCategories, err := r.repo.MenuCategories(ctx)
	r.Assert().NoError(err)
	r.Assert().NotEmpty(menuCategories)

	var argsInsertMenu = &entities.Menu{
		MenuCategoriesID: argsInsert.ID,
		Name:             "Kopi Susu Gula Aren",
		Description:      "this product is combination between one shoot espresso and palm sugar",
		Price:            25000,
		CreatedAt:        time.Now(),
	}

	err = r.repo.InsertMenu(ctx, argsInsertMenu)
	r.Assert().NoError(err)

	detailMenu, err := r.repo.DetailMenu(ctx, argsInsertMenu.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(argsInsertMenu.Name, detailMenu.Name)
	r.Assert().Equal(argsInsertMenu.Description, detailMenu.Description)

	var dateMenu = time.Now()

	var argsUpdateMenu = entities.UpdateMenu{
		ID:          argsInsertMenu.ID,
		Name:        "Kopi Susu Gula Aren Strong",
		Description: "this product is combination between two shoot espresso and palm sugar",
		Price:       30000,
		UpdatedAt:   &dateMenu,
	}

	err = r.repo.UpdateMenuByID(ctx, argsUpdateMenu)
	r.Assert().NoError(err)

	detailMenuAfterUpdate, err := r.repo.DetailMenu(ctx, argsInsertMenu.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(argsUpdateMenu.Name, detailMenuAfterUpdate.Name)
	r.Assert().Equal(argsUpdateMenu.Description, detailMenuAfterUpdate.Description)

	listMenu, err := r.repo.ListMenu(ctx)
	r.Assert().NoError(err)
	r.Assert().NotEmpty(listMenu)
}
