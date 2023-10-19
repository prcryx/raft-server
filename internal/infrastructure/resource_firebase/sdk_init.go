package resource_firebase

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go"
	"github.com/prcryx/raft-server/config"
	e "github.com/prcryx/raft-server/internal/common/err"
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
