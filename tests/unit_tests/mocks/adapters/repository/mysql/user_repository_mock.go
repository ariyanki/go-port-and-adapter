package mysql

import (
	"errors"
	"go-port-and-adapter/ports/repository/constants/user_status"
	"go-port-and-adapter/ports/repository/dto"

	"github.com/stretchr/testify/mock"
)

type (
	userRepository struct {
		mock.Mock
	}
)

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (db *userRepository) GetUserByUsername(username string, user *dto.UserDto) error {

	db.Called(username)
	if username == "admin@admin.com" {
		user.Status = user_status.Active
		user.Password = "26fce84bdf52dced60502b2f126140d3edfbd5b0d7c6586bb79b5cb45680debf"
		user.PasswordSalt = "e02e8140-06a1-45ce-9240-b610ee58add4"

		return nil
	}
	return nil
}

func (db *userRepository) CreateUser(user dto.CreateUserDto) error {
	db.Called(user)
	return nil
}

func (db *userRepository) GetUserById(id int, user *dto.UserDto) error {
	db.Called(id)
	if id == 123 {
		return errors.New("Not Found")
	}

	user.Status = user_status.Active
	user.Password = "26fce84bdf52dced60502b2f126140d3edfbd5b0d7c6586bb79b5cb45680debf"
	user.PasswordSalt = "e02e8140-06a1-45ce-9240-b610ee58add4"

	return nil
}

func (db *userRepository) UpdateUser(user dto.UpdateUserDto) error {
	db.Called(user)
	return nil
}

func (db *userRepository) DeleteUser(id int) error {
	db.Called(id)
	return nil
}

func (db *userRepository) ListUser(listRequest dto.ListRequestDto, users *[]dto.UserDto) error {
	db.Called(listRequest)
	return nil

}
