package dtos

import "time"

type Users struct {
	ID        int64      `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Gender    string     `json:"gender"`
	Birthday  string     `json:"birthday"`
	Points    int64      `json:"points"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type UpdateUserProfileRequest struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}

type UpdateUserPointRequest struct {
	ID     int64 `json:"id"`
	Points int64 `json:"points"`
}

type UserSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	Token string `json:"token"`
}
