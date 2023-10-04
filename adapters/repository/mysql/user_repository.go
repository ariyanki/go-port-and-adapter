package mysql

import (
	"go-port-and-adapter/ports/repository/dto"

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
	db.Where("id = ?", ID).Find(&user)
	return user, nil
}

func (db *UserRepository) GetUserByUsername(username string) (dto.UserDto, error) {
	var user dto.UserDto
	db.Where("username = ?", username).Find(&user)
	return user, nil
}
