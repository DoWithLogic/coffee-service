package entities

import (
	"time"

	"github.com/DoWithLogic/coffee-service/internal/users/dtos"
	"github.com/DoWithLogic/coffee-service/pkg/utils"
)

type (
	Gender   string
	Password string
)

const (
	Male   Gender = "mele"
	Femele Gender = "femele"
	Other  Gender = "other"
)

func (x Gender) String() string {
	return string(x)
}

func (x Password) String() string {
	return string(x)
}

func (x Password) Encrypt() *string {
	return utils.HashPassword(string(x))
}

func (x Password) VerifyPassword(encryptedPassword string) bool {
	return utils.VerifyPassword(string(x), encryptedPassword)
}

type Users struct {
	ID        int64      `db:"id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  Password   `db:"password"`
	Gender    Gender     `db:"gender"`
	Birthday  string     `db:"birthday"`
	Points    int64      `db:"points"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type UpdateUserProfile struct {
	ID        int64      `db:"id"`
	Username  string     `db:"username"`
	Email     string     `db:"email"`
	Password  Password   `db:"password"`
	Gender    Gender     `db:"gender"`
	Birthday  string     `db:"birthday"`
	UpdatedAt *time.Time `db:"updated_at"`
}

type UpdateUserPoint struct {
	ID        int64      `db:"id"`
	Points    int64      `db:"points"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func NewUsers(req dtos.Users) Users {
	return Users{
		Username:  req.Username,
		Email:     req.Email,
		Password:  Password(*Password(req.Password).Encrypt()),
		Gender:    Gender(req.Gender),
		Birthday:  req.Birthday,
		CreatedAt: time.Now(),
	}
}

func NewDetailResponse(req Users) dtos.Users {
	return dtos.Users{
		ID:        req.ID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  req.Password.String(),
		Gender:    req.Gender.String(),
		Birthday:  req.Birthday,
		Points:    req.Points,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}

func NewUpdateUserProfile(req dtos.UpdateUserProfileRequest) UpdateUserProfile {
	var (
		now = time.Now()
		pwd = *Password(req.Password).Encrypt()
	)

	return UpdateUserProfile{
		ID:        req.ID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  Password(pwd),
		Gender:    Gender(req.Gender),
		Birthday:  req.Birthday,
		UpdatedAt: &now,
	}
}

func NewUpdateUserPoint(req dtos.UpdateUserPointRequest) UpdateUserPoint {
	var now = time.Now()
	return UpdateUserPoint{
		ID:        req.ID,
		Points:    req.Points,
		UpdatedAt: &now,
	}
}
