package datasoruces

import (
	"context"

	"firebase.google.com/go/auth"
)

type IAuthDataSource interface {
	SignUpWithEmailAndPassword(string, string) (*auth.UserRecord, error)
}

type AuthDataSource struct {
	authClient *auth.Client
}

var _ IAuthDataSource = (*AuthDataSource)(nil)

func NewAuthDataSource(authClient *auth.Client) (*AuthDataSource, error) {
	return &AuthDataSource{authClient: authClient}, nil
}

func (authDataSoruce *AuthDataSource) SignUpWithEmailAndPassword(email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	user, err := authDataSoruce.authClient.CreateUser(context.Background(), params)
	if err != nil {

		return nil, err
	}

	return user, nil
}
