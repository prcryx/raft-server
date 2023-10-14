package auth

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/prcryx/raft-server/internal/apis/auth/dto"
)


func AuthenticateUser(authDto *dto.AuthDTO, ac *AuthController) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(authDto.Email).
		Password(authDto.Password)

	user, err := ac.AuthClient.CreateUser(context.Background(), params)
	if err != nil {

		return nil, err
	}

	return user, nil
}
