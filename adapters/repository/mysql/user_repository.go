package mysql

import (

	"go-port-and-adapter/ports/repository/dto"
	repoError "go-port-and-adapter/ports/repository/constants/error"

	"github.com/jinzhu/gorm"
)

type (
	UserRepository struct {
		*gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db.Table("app_users"),
	}
}

func (db *UserRepository) CreateUser(user dto.UserDto) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (db *UserRepository) GetUserByID(ID string) (dto.UserDto, error) {
	var user dto.UserDto
	if db.Where("id = ?", ID).Find(&user).RecordNotFound() {
		return user, repoError.RecordNotFound
	}
	return user, nil
}

func (db *UserRepository) GetUserByUsername(username string) (dto.UserDto, error) {
	var user dto.UserDto
	if db.Where("username = ?", username).Find(&user).RecordNotFound() {
		return user, repoError.RecordNotFound
	}
	return user, nil
}
