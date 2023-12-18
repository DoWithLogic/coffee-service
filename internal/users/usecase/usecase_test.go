package usecase_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DoWithLogic/coffee-service/config"
	"github.com/DoWithLogic/coffee-service/internal/users/dtos"
	"github.com/DoWithLogic/coffee-service/internal/users/entities"
	"github.com/DoWithLogic/coffee-service/internal/users/mocks"
	"github.com/DoWithLogic/coffee-service/internal/users/usecase"
	"github.com/DoWithLogic/coffee-service/pkg/apperrors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

type eqUsersMatcher struct {
	users entities.Users
}

func insertUsersMatcher(u entities.Users) gomock.Matcher {
	return &eqUsersMatcher{
		users: u,
	}
}

func (e eqUsersMatcher) Matches(x interface{}) bool {
	arg, ok := x.(*entities.Users)
	if !ok {
		return false
	}

	return arg.Username == e.users.Username && arg.Email == e.users.Email && arg.Birthday == e.users.Birthday
}

func (e eqUsersMatcher) String() string {
	return fmt.Sprintf("%v", e.users.Username)
}

func Test_usecase_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		repo = mocks.NewMockRepositories(ctrl)
		uc   = usecase.NewUsecase(repo, config.Config{})
	)

	var request = dtos.Users{
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.PASSWORD,
		Gender:   "mele",
		Birthday: "1997-01-25",
	}

	t.Run("positive_signup_ok", func(t *testing.T) {
		repo.EXPECT().InsertUsers(context.Background(), insertUsersMatcher(entities.NewUsers(request))).Return(nil)

		err := uc.SignUp(context.Background(), &request)
		require.NoError(t, err)
	})

	t.Run("negative_signup_failed_insert", func(t *testing.T) {
		repo.EXPECT().InsertUsers(context.Background(), insertUsersMatcher(entities.NewUsers(request))).Return(sql.ErrNoRows)

		err := uc.SignUp(context.Background(), &request)
		require.ErrorIs(t, err, sql.ErrNoRows)
	})
}

func Test_usecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		repo = mocks.NewMockRepositories(ctrl)
		uc   = usecase.NewUsecase(repo, config.Config{})
	)

	email := faker.Email()
	pwd := "testing"

	var userDetails = entities.Users{
		ID:        1,
		Username:  faker.Username(),
		Email:     email,
		Password:  entities.Password(*entities.Password(pwd).Encrypt()),
		Gender:    entities.Male,
		Birthday:  "1997-01-25",
		Points:    0,
		CreatedAt: time.Now(),
	}

	t.Run("positive_SignIn_Ok", func(t *testing.T) {
		repo.EXPECT().UserDetailByEmail(context.Background(), email).Return(userDetails, nil)

		response, err := uc.SignIn(context.Background(), dtos.UserSignInRequest{Email: email, Password: pwd})
		require.NoError(t, err)
		require.NotEmpty(t, response.Token)
	})

	t.Run("negative_SignIn_ErrNotFound", func(t *testing.T) {
		repo.EXPECT().UserDetailByEmail(context.Background(), email).Return(entities.Users{}, sql.ErrNoRows)

		response, err := uc.SignIn(context.Background(), dtos.UserSignInRequest{Email: email, Password: pwd})
		require.ErrorIs(t, err, apperrors.ErrNotFound)
		require.Equal(t, response, dtos.UserSignInResponse{})
	})

	t.Run("negative_SignIn_FailedGetDetail", func(t *testing.T) {
		repo.EXPECT().UserDetailByEmail(context.Background(), email).Return(entities.Users{}, apperrors.ErrInternalServer)

		response, err := uc.SignIn(context.Background(), dtos.UserSignInRequest{Email: email, Password: pwd})
		require.ErrorIs(t, err, apperrors.ErrInternalServer)
		require.Equal(t, response, dtos.UserSignInResponse{})
	})

	t.Run("negative_SignIn_InvalidPassword", func(t *testing.T) {
		repo.EXPECT().UserDetailByEmail(context.Background(), email).Return(userDetails, nil)

		response, err := uc.SignIn(context.Background(), dtos.UserSignInRequest{Email: email, Password: "invalid_password"})
		require.ErrorIs(t, err, apperrors.ErrInvalidPassword)
		require.Equal(t, response, dtos.UserSignInResponse{})
	})
}

func Test_usecase_UserDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var (
		repo = mocks.NewMockRepositories(ctrl)
		uc   = usecase.NewUsecase(repo, config.Config{})
	)

	var userDetails = entities.Users{
		ID:        1,
		Username:  faker.Username(),
		Email:     faker.Email(),
		Password:  entities.Password(*entities.Password(faker.PASSWORD).Encrypt()),
		Gender:    entities.Male,
		Birthday:  "1997-01-25",
		Points:    0,
		CreatedAt: time.Now(),
	}

	t.Run("positive_UserDetail_Ok", func(t *testing.T) {
		repo.EXPECT().UserDetail(context.Background(), userDetails.ID).Return(userDetails, nil)

		userDetail, err := uc.UserDetail(context.Background(), userDetails.ID)
		require.NoError(t, err)
		require.Equal(t, userDetail, entities.NewDetailResponse(userDetails))
	})
}
