package repository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/users"
	"github.com/DoWithLogic/coffee-service/internal/users/entities"
	"github.com/DoWithLogic/coffee-service/internal/users/repository"
	"github.com/DoWithLogic/coffee-service/pkg/databases"
	"github.com/go-faker/faker/v4"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
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
	repo users.Repositories
	suite.Suite
}

func TestSuiteRepository(t *testing.T) {
	defer func() {
		// err := migrationDOWN(MIGRATION_PATH, db.DB)
		// require.NoError(t, err)
	}()

	suite.Run(t, &repositoryTestSuite{repo: repository.NewRepository(db)})
}

func (r *repositoryTestSuite) Test_all_repositories_function() {
	var users = &entities.Users{
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  entities.Password(faker.Password()),
		Gender:    entities.Male,
		Birthday:  "1997-01-25",
		CreatedAt: time.Now(),
	}

	err := r.repo.InsertUsers(context.Background(), users)
	r.Assert().NoError(err)
	r.Assert().NotEqual(0, users.ID)

	detail, err := r.repo.UserDetail(context.Background(), users.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(users.ID, detail.ID)

	updatePwd := *entities.Password(faker.Word()).Encrypt()
	now := time.Now()

	updateUserProfile := entities.UpdateUserProfile{
		ID:        detail.ID,
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  entities.Password(updatePwd),
		Gender:    entities.Femele,
		Birthday:  "1997-01-01",
		UpdatedAt: &now,
	}

	err = r.repo.UpdateUserProfile(context.Background(), updateUserProfile)
	r.Assert().NoError(err)

	detailAfterProfile, err := r.repo.UserDetail(context.Background(), detail.ID)
	r.Assert().NoError(err)
	r.Assert().Equal(updateUserProfile.Username, detailAfterProfile.Username)
	r.Assert().Equal(updateUserProfile.Email, detailAfterProfile.Email)

	var updateUserPointRequest = entities.UpdateUserPoint{
		ID:        users.ID,
		Points:    500,
		UpdatedAt: &now,
	}
	err = r.repo.UpdateUserPoint(context.Background(), updateUserPointRequest)
	r.Assert().NoError(err)

	detailByEmail, err := r.repo.UserDetailByEmail(context.Background(), updateUserProfile.Email)
	r.Assert().NoError(err)
	r.Assert().Equal(detailByEmail.Points, updateUserPointRequest.Points)
}
