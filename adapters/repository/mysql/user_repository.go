package mysql

import (
	"go-port-and-adapter/ports/repository/dto"

	"github.com/jinzhu/gorm"
)

type (
	userRepository struct {
		*gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db.Table("app_users"),
	}
}

func (db *userRepository) GetUserByUsername(username string, user *dto.UserDto) error {
	if err := db.Where("username = ?", username).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (db *userRepository) CreateUser(user dto.CreateUserDto) error {
	user.Status = 1
	user.UpdatedBy = user.CreatedBy
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (db *userRepository) GetUserById(id int, user *dto.UserDto) error {
	user.UpdatedBy = user.CreatedBy
	if err := db.Where("id = ?", id).Find(&user).Error; err != nil {
		return err
	}
	return nil
}

func (db *userRepository) UpdateUser(user dto.UpdateUserDto) error {
	if err := db.Model(&dto.UserDto{}).Where("id = ?", user.ID).Update(user).Error; err != nil {
		return err
	}
	return nil
}

func (db *userRepository) DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(&dto.UserDto{}).Error; err != nil {
		return err
	}
	return nil
}

func (db *userRepository) ListUser(listRequest dto.ListRequestDto, users *[]dto.UserDto) error {
	query := db.DB
	// filter
	for field, value := range listRequest.Filter {
		query = query.Where(field + " LIKE " + "'%" + value + "%'")
	}
	// sort
	for field, sort := range listRequest.OrderBy {
		query = query.Order(field + " " + sort)
	}
	// paging
	if listRequest.PageSize > 0 && listRequest.Page > 0 {
		query = query.Limit(listRequest.PageSize).Offset((listRequest.Page - 1) * listRequest.PageSize)
	}
	// Find
	if err := query.Find(users).Error; err != nil {
		return err
	}
	return nil

}
