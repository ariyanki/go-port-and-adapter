package storage

import (
	"encoding/base64"
	"go-port-and-adapter/ports/storage/dto"
	"os"

	"github.com/spf13/viper"
)

type (
	storageManager struct{}
)

func NewStorageManager() *storageManager {
	return &storageManager{}
}

func (db *storageManager) StoreUserPhoto(userPhoto dto.UserPhotoDto) error {
	filedir := viper.GetString("filedir")

	// Decode the base64-encoded file data
	decodedData, err := base64.StdEncoding.DecodeString(userPhoto.PhotoFiledata)
	if err != nil {
		return err
	}

	// Create a destination file with the provided file name
	dst, err := os.Create(filedir + "/user_photo/" + userPhoto.PhotoFilename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Write the decoded file data to the destination file
	_, err = dst.Write(decodedData)
	if err != nil {
		return err
	}
	return nil
}
