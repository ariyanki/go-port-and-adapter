
package dto

type (
	SignInRequest struct {
		Username      string    `json:"username"`
		Password     string    `json:"password"`
	}

	SignInResponse struct {
		Token      string    `json:"token"`
	}
)