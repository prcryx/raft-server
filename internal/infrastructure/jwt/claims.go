package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UID     uint   `json:"uid"`
	PhoneNo string `json:"phoneNo"`
	jwt.RegisteredClaims
}

func NewCustomClaims(uid uint, phoneNo string) CustomClaims {
	return CustomClaims{
		uid,
		phoneNo,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
