package auth

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

type AuthController struct {
	AuthClient *auth.Client
}

func NewAuthController(app *firebase.App) (*AuthController, error) {
	authClient, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return &AuthController{AuthClient: authClient}, nil
}

func (ac *AuthController) SignUpWithEmailAndPassword(email, password string) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)

	user, err := ac.AuthClient.CreateUser(context.Background(), params)
	if err != nil {
		log.Printf("Error creating user: %v\n", err)
		return nil, err
	}

	return user, nil
}