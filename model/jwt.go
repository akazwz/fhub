package model

import (
	"github.com/golang-jwt/jwt"
)

type MyCustomClaims struct {
	ID         string `json:"id"`
	BufferTime int64  `json:"buffer_time"`
	jwt.StandardClaims
}
