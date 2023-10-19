package dto

import "time"

type (
	CreateUserDto struct {
		Username      string    `json:"username"`
		Fullname      string    `json:"fullname"`
		Phonenumber   string    `json:"phonenumber"`
		Email         string    `json:"email"`
		PhotoFilename string    `json:"photo_filename"`
		Birthdate     time.Time `json:"birthdate"`
		Gender        string    `json:"gender"`
		City          string    `json:"city"`
		Address       string    `json:"address"`
		Status        int       `json:"status"`
		CreatedBy     int       `json:"created_by"`
		UpdatedBy     int       `json:"updated_by"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
	}

	UpdateUserDto struct {
		ID            int       `json:"id"`
		Username      string    `json:"username"`
		Fullname      string    `json:"fullname"`
		Phonenumber   string    `json:"phonenumber"`
		Email         string    `json:"email"`
		PhotoFilename string    `json:"photo_filename"`
		Birthdate     time.Time `json:"birthdate"`
		Gender        string    `json:"gender"`
		City          string    `json:"city"`
		Address       string    `json:"address"`
		Status        int       `json:"status"`
		UpdatedBy     int       `json:"updated_by"`
		UpdatedAt     time.Time `json:"updated_at"`
	}

	UserDto struct {
		ID                int       `json:"id"`
		FacebookLoginId   string    `json:"facebook_login_id"`
		GoogleLoginId     string    `json:"google_login_id"`
		Username          string    `json:"username"`
		Password          string    `json:"password"`
		PasswordSalt      string    `json:"password_salt"`
		Fullname          string    `json:"fullname"`
		Phonenumber       string    `json:"phonenumber"`
		Email             string    `json:"email"`
		Longitude         string    `json:"longitude"`
		Latitude          string    `json:"latitude"`
		PhotoFilename     string    `json:"photo_filename"`
		IsLoggedin        int       `json:"is_loggedin"`
		LoginAttempt      int       `json:"login_attempt"`
		LastLoggedinAt    time.Time `json:"last_loggedin_at"`
		AccessSource      string    `json:"access_source"`
		LastAccessAt      time.Time `json:"last_access_at"`
		GoogleLoginData   string    `json:"google_login_data"`
		FacebookLoginData string    `json:"facebook_login_data"`
		Birthdate         time.Time `json:"birthdate"`
		Gender            string    `json:"gender"`
		FcmToken          string    `json:"fcm_token"`
		City              string    `json:"city"`
		Address           string    `json:"address"`
		MerchantId        int       `json:"merchant_id"`
		AppleLoginId      string    `json:"apple_login_id"`
		AppleLoginData    string    `json:"apple_login_data"`
		Status            int       `json:"status"`
		CreatedBy         int       `json:"created_by"`
		UpdatedBy         int       `json:"updated_by"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		DeletedAt         time.Time `json:"deleted_at"`
	}
)
