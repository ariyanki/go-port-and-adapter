package dto

import jwt "github.com/dgrijalva/jwt-go"

//JwtCustomClaims JwtCustomClaims
type JwtCustomClaims struct {
	ID       int    `json:"id"`
	jwt.StandardClaims
}