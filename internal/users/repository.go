package users

import (
	"context"

	"github.com/DoWithLogic/coffee-service/internal/users/entities"
)

type Repositories interface {
	InsertUsers(ctx context.Context, users *entities.Users) error
	UserDetail(ctx context.Context, userID int64) (users entities.Users, err error)
	UserDetailByEmail(ctx context.Context, Email string) (users entities.Users, err error)
	UpdateUserProfile(ctx context.Context, users entities.UpdateUserProfile) error
	UpdateUserPoint(ctx context.Context, request entities.UpdateUserPoint) error
}
