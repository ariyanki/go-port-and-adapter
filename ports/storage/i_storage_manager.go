package storage

import "go-port-and-adapter/ports/storage/dto"

// IStorageManager is an interface to store file in storage
type IStorageManager interface {

	//StoreUserPhoto
	StoreUserPhoto(userPhoto dto.UserPhotoDto) error
}
