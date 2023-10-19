// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	auth2 "github.com/prcryx/raft-server/internal/application/apis/auth"
	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/data/datasoruces"
	"github.com/prcryx/raft-server/internal/data/repository_impl"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/domain/usecases"
	"github.com/prcryx/raft-server/internal/infrastructure/resource_firebase"
)

// Injectors from wire.go:

func InitFirebaseApp(envConfig *config.EnvConfig) (*firebase.App, error) {
	app, err := resource_firebase.InitFirebaseApp(envConfig)
	if err != nil {
		return nil, err
	}
	return app, nil
}

func InitFirebaseAuthClient(app *firebase.App) (*auth.Client, error) {
	client, err := resource_firebase.SetupFirebaseAuth(app)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InitializeControllerRegistry(authClient *auth.Client) (*container.ControllerRegistry, error) {
	authDataSource, err := datasoruces.NewAuthDataSource(authClient)
	if err != nil {
		return nil, err
	}
	authRepositoryImpl := repository_impl.NewAuthRepositoryImpl(authDataSource)
	authUseCase := usecases.NewAuthUseCase(authRepositoryImpl)
	authController := auth2.NewAuthController(authUseCase)
	controllerRegistry := container.NewControllerRegistry(authController)
	return controllerRegistry, nil
}

func InitServer(controllerRegistry *container.ControllerRegistry, configVars *config.EnvConfig, version string) (*types.Server, error) {
	appApp := app.NewApp(controllerRegistry, configVars, version)
	typesServer, err := server.NewServer(appApp)
	if err != nil {
		return nil, err
	}
	return typesServer, nil
}
