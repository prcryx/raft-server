//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/infrastructure/postgres"
	"github.com/prcryx/raft-server/internal/infrastructure/twilio"
	"gorm.io/gorm"
)

// init twilio app
func InitTwilioApp(conf *config.EnvConfig) (*twilio.TwilioApp, error) {
	wire.Build(
		twilio.NewTwilioApp,
	)
	return nil, nil
}

// init Database
func InitDatabase(envConfig *config.EnvConfig) (*gorm.DB, error) {
	wire.Build(
		postgres.CreatePostgresDatabase,
	)
	return nil, nil
}

// init ControllerRegistry

func InitializeControllerRegistry(db *gorm.DB, twilioApp *twilio.TwilioApp, conf *config.EnvConfig) (*container.ControllerRegistry, error) {
	wire.Build(
		container.NewControllerRegistry,
		DataSourceSet,
		RepositorySet,
		UseCaseSet,
		ControllerSet,
		OtherServices,
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
