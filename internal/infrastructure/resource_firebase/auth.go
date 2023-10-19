package resource_firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

// setup firebase auth
func SetupFirebaseAuth(app *firebase.App) (*auth.Client, error) {
	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return auth, nil
}
