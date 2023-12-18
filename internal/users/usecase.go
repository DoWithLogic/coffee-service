package users

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/users/dtos"
)

type UseCases interface {
	SignUp(ctx context.Context, request *dtos.Users) error
	SignIn(ctx context.Context, request dtos.UserSignInRequest) (response dtos.UserSignInResponse, err error)
	UserDetail(ctx context.Context, id int64) (dtos.Users, error)
}
