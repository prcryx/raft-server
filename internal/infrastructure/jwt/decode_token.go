package jwt

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	e "github.com/prcryx/raft-server/internal/common/err"
)

func DecodeToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.UnacceptedSigningError()
		}
		return []byte(""), nil
	})

	fmt.Println(token.Claims)

	if err != nil {
		log.Println(err)
		return nil, e.UnauthorizedException()
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Println(err)
		return nil, e.UnauthorizedException()
	}
}
