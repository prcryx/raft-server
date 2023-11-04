package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/prcryx/raft-server/config"
)

type IJwtStrategy interface {
	GenerateToken(uint, string) string
}

type JwtStrategy struct {
	secretKey []byte
}

func NewJwtStrategy(conf *config.EnvConfig) *JwtStrategy {
	return &JwtStrategy{
		secretKey: []byte(conf.Server.JWT.AccessTokenSecret),
	}
}

var _ IJwtStrategy = (*JwtStrategy)(nil)

func (jwtStrategy *JwtStrategy) GenerateToken(uid uint, phoneNo string) string {
	claims := NewCustomClaims(uid, phoneNo)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtStrategy.secretKey)
	if err != nil {
		return ""
	}
	return tokenString
}
