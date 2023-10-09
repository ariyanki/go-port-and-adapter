
package dto

import jwt "github.com/dgrijalva/jwt-go"

type (
	SignInRequest struct {
		Username      string    `json:"username"`
		Password     string    `json:"password"`
	}

	SignInResponse struct {
		Token      string    `json:"token"`
	}

	//JwtCustomClaims JwtCustomClaims
	JwtCustomClaims struct {
		ID       int    `json:"id"`
		jwt.StandardClaims
	}
)
