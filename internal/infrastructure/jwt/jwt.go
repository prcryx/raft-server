package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prcryx/raft-server/config"
	e "github.com/prcryx/raft-server/internal/common/err"
)

type IJwtStrategy interface {
	GenerateToken(uint, string) string
	DecodeToken(string) (*jwt.MapClaims, error)
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

func (jwtStrategy *JwtStrategy) DecodeToken(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		fmt.Println(t)
		fmt.Println(t.Method)
		fmt.Println(t.Signature)
		fmt.Println(t.Header["alg"])
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, e.UnacceptedSigningError()
		}
		return jwtStrategy.secretKey, nil
	})

	fmt.Println("token.Valid: ", token.Valid)

	if err != nil {
		return nil, e.UnauthorizedException()
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// ctx := context.WithValue(r.Context(), "props", claims)
		// Access context values in handlers like this
		// props, _ := r.Context().Value("props").(jwt.MapClaims)
		// next.ServeHTTP(w, r.WithContext(ctx))
		return &claims, nil
	} else {
		// fmt.Println(err)
		return nil, e.UnauthorizedException()
	}
}
