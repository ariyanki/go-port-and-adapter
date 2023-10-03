package dto

import "time"

type (
	CreateUserDto struct {
		ID        int       `json:"id"`
		Username      string    `json:"username"`
		Password     string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	UserDto struct {
		ID        int       `json:"id"`
		Username      string    `json:"username"`
		Password     string    `json:"password"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		DeletedAt time.Time `json:"deleted_at"`
	}
)