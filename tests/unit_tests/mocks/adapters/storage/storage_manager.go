package storage

import "go-port-and-adapter/ports/storage/dto"

type (
	storageManager struct{}
)

func NewStorageManager() *storageManager {
	return &storageManager{}
}

func (db *storageManager) StoreUserPhoto(userPhoto dto.UserPhotoDto) error {
	return nil
}
