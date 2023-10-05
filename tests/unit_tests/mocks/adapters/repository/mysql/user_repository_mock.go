package mysql

import (
	"go-port-and-adapter/ports/repository/dto"
	"go-port-and-adapter/ports/repository/constants/user_status"

	"github.com/stretchr/testify/mock"
)

type (
	UserRepository struct {
		mock.Mock
	}
)

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (db *UserRepository) CreateUser(user dto.UserDto) error {
	db.Called(user)
	return nil
}

func (db *UserRepository) GetUserByID(ID string) (dto.UserDto, error) {
	db.Called(ID)
	return dto.UserDto{}, nil
}

func (db *UserRepository) GetUserByUsername(username string) (dto.UserDto, error) {
	db.Called(username)
	if username == "admin@admin.com" {
		return dto.UserDto{
			Status: user_status.Active,
			Password: "26fce84bdf52dced60502b2f126140d3edfbd5b0d7c6586bb79b5cb45680debf",
			PasswordSalt: "e02e8140-06a1-45ce-9240-b610ee58add4",

		}, nil
	}
	return dto.UserDto{}, nil
}
