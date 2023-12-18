package repository

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/users"
	"github.com/DoWithLogic/coffee-service/internal/users/entities"
	"github.com/DoWithLogic/coffee-service/pkg/observability"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type repo struct {
	db  *sqlx.DB
	log *zerolog.Logger
}

func NewRepository(db *sqlx.DB) users.Repositories {
	return &repo{db: db, log: observability.NewZeroLogHook().Z()}
}

func (r *repo) InsertUsers(ctx context.Context, users *entities.Users) error {
	sqlResult, err := r.db.NamedExecContext(ctx, insertUsers, &users)
	if err != nil {
		r.log.Err(err).Msg("[users][InsertUsers]NamedExecContext")

		return err
	}

	users.ID, err = sqlResult.LastInsertId()
	if err != nil {
		r.log.Err(err).Msg("[users][InsertUsers]LastInsertId")

		return err
	}

	return nil
}

func (r *repo) UserDetail(ctx context.Context, userID int64) (users entities.Users, err error) {
	if err := r.db.GetContext(ctx, &users, userDetail, userID); err != nil {
		r.log.Err(err).Msg("[users][UserDetail]GetContext")

		return users, err
	}

	return users, nil
}

func (r *repo) UserDetailByEmail(ctx context.Context, Email string) (users entities.Users, err error) {
	if err := r.db.GetContext(ctx, &users, userDetailByEmail, Email); err != nil {
		r.log.Err(err).Msg("[users][UserDetailByEmail]GetContext")

		return users, err
	}

	return users, nil
}

func (r *repo) UpdateUserProfile(ctx context.Context, users entities.UpdateUserProfile) error {
	if _, err := r.db.NamedExecContext(ctx, updateUserProfile, users); err != nil {
		r.log.Err(err).Msg("[users][UpdateUserProfile]NamedExecContext")

		return err
	}

	return nil
}

func (r *repo) UpdateUserPoint(ctx context.Context, request entities.UpdateUserPoint) error {
	if _, err := r.db.ExecContext(ctx, updateUserPoint, request.Points, request.UpdatedAt, request.ID); err != nil {
		r.log.Err(err).Msg("[users][UpdateUserPoint]ExecContext")

		return err
	}

	return nil
}
