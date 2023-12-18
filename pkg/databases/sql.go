package databases

import (
	"context"
	"fmt"

	"github.com/DoWithLogic/coffee-service/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMySQLDB(ctx context.Context, cfg config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Asia%%2FJakarta", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil
}
