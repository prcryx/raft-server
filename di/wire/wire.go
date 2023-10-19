//go:build wireinject

package wire

import (
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/google/wire"
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/domain/types"
	rf "github.com/prcryx/raft-server/internal/infrastructure/resource_firebase"
)

//init firebase app

func InitFirebaseApp(envConfig *config.EnvConfig) (*firebase.App, error) {
	wire.Build(
		rf.InitFirebaseApp,
	)
	return nil, nil
}

//init auth client

func InitFirebaseAuthClient(app *firebase.App) (*auth.Client, error) {
	wire.Build(
		rf.SetupFirebaseAuth,
	)
	return nil, nil
}

// init ControllerRegistry

func InitializeControllerRegistry(authClient *auth.Client) (*container.ControllerRegistry, error) {
	wire.Build(
		container.NewControllerRegistry,
		DataSourceSet,
		RepositorySet,
		UseCaseSet,
		ControllerSet,
	)

	return nil, nil
}

//Intialize Server

func InitServer(controllerRegistry *container.ControllerRegistry, configVars *config.EnvConfig, version string) (*types.Server, error) {
	wire.Build(
		server.NewServer,
		app.NewApp,
	)

	return nil, nil
}
