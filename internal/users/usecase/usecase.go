package usecase

import (
	"context"
	"database/sql"

	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/users"
	"github.com/DoWithLogic/coffee-service/internal/users/dtos"
	"github.com/DoWithLogic/coffee-service/internal/users/entities"
	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/DoWithLogic/coffee-service/pkg/auth"
	"github.com/DoWithLogic/coffee-service/pkg/observability"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type usecase struct {
	repo users.Repositories
	cfg  config.Config
	log  *zerolog.Logger
}

func NewUsecase(r users.Repositories, cfg config.Config) users.UseCases {
	return &usecase{repo: r, cfg: cfg, log: observability.NewZeroLogHook().Z()}
}
func (uc *usecase) SignUp(ctx context.Context, request *dtos.Users) error {
	usersEntity := entities.NewUsers(*request)
	if err := uc.repo.InsertUsers(ctx, &usersEntity); err != nil {
		return err
	}

	request.ID = usersEntity.ID

	return nil
}

func (uc *usecase) SignIn(ctx context.Context, request dtos.UserSignInRequest) (response dtos.UserSignInResponse, err error) {
	userDetail, err := uc.repo.UserDetailByEmail(ctx, request.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return response, apperrors.ErrNotFound
		}

		return response, err
	}

	if ok := entities.Password(request.Password).VerifyPassword(userDetail.Password.String()); !ok {
		uc.log.Err(apperrors.ErrInvalidPassword).Msg("[users][SignIn]VerifyPassword")

		return response, apperrors.ErrInvalidPassword
	}

	var secretKey = auth.Authorization(uc.cfg.Authorization.SecretKey)
	response.Token, err = secretKey.GenerateToken(userDetail.Email)
	if err != nil {
		uc.log.Err(apperrors.ErrInvalidPassword).Msg("[users][SignIn]GenerateToken")

		return response, err
	}

	return response, err
}
func (uc *usecase) UserDetail(ctx context.Context, id int64) (response dtos.Users, err error) {
	users, err := uc.repo.UserDetail(ctx, id)
	if err != nil {
		return response, err
	}

	return entities.NewDetailResponse(users), nil
}
