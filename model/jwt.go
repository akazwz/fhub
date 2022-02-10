package model

import (
	"github.com/golang-jwt/jwt"
	uuid "github.com/satori/go.uuid"
)

type MyCustomClaims struct {
	UID        uuid.UUID `json:"uid"`
	Username   string    `json:"username"`
	BufferTime int64     `json:"buffer_time"`
	jwt.StandardClaims
}
