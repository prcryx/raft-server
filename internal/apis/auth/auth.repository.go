package auth

import (
	"context"

	"firebase.google.com/go/auth"
	"github.com/prcryx/raft-server/internal/apis/auth/dto"
)

type AuthRepository struct {
	ac      *AuthController
	authDto *dto.AuthDTO
}

func (repo *AuthRepository) AuthenticateUser() (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(repo.authDto.Email).
		Password(repo.authDto.Password)

	user, err := repo.ac.AuthClient.CreateUser(context.Background(), params)
	if err != nil {

		return nil, err
	}

	return user, nil
}
