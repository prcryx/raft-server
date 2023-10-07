package db

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	e "github.com/prcryx/raft-server/common/err"
	"github.com/prcryx/raft-server/config"
	"google.golang.org/api/option"
)

// Firebase admin SDK initialization
func InitFirebaseApp(env *config.EnvConfig) (*firebase.App, error) {
	serviceAccountKeyFilePath, err := filepath.Abs(env.ServiceAccountKeyFile)
	if err != nil {
		log.Fatal(e.UnableToLoadServiceJson)
	}
	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf(e.FirebaseLoadError)
		return nil, err
	}
	return app, nil
}

// setup firebase auth
func SetupFirebaseAuth(app *firebase.App) *auth.Client {
	auth, err := app.Auth(context.Background())
	if err != nil {
		log.Fatal(e.FirebaseLoadError)
	}
	return auth
}
