// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	"github.com/prcryx/raft-server/internal/application/apis/auth"
	"github.com/prcryx/raft-server/internal/application/apis/feeds"
	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/server"
	"github.com/prcryx/raft-server/internal/data/datasoruces"
	"github.com/prcryx/raft-server/internal/data/repository_impl"
	"github.com/prcryx/raft-server/internal/domain/types"
	"github.com/prcryx/raft-server/internal/domain/usecases"
	"github.com/prcryx/raft-server/internal/infrastructure/jwt"
	"github.com/prcryx/raft-server/internal/infrastructure/postgres"
	"github.com/prcryx/raft-server/internal/infrastructure/twilio"
	"gorm.io/gorm"
)

// Injectors from wire.go:

// init twilio app
func InitTwilioApp(conf *config.EnvConfig) (*twilio.TwilioApp, error) {
	twilioApp := twilio.NewTwilioApp(conf)
	return twilioApp, nil
}

// init Database
func InitDatabase(envConfig *config.EnvConfig) (*gorm.DB, error) {
	db, err := postgres.CreatePostgresDatabase(envConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitializeControllerRegistry(db *gorm.DB, twilioApp *twilio.TwilioApp, conf *config.EnvConfig) (*container.ControllerRegistry, error) {
	jwtStrategy := jwt.NewJwtStrategy(conf)
	authDataSource, err := datasoruces.NewAuthDataSource(db, twilioApp, jwtStrategy)
	if err != nil {
		return nil, err
	}
	authRepositoryImpl := repository_impl.NewAuthRepositoryImpl(authDataSource)
	authUseCase := usecases.NewAuthUseCase(authRepositoryImpl)
	authController := auth.NewAuthController(authUseCase)
	feedUseCase := usecases.NewFeedUseCase()
	feedController := feed.NewFeedController(feedUseCase)
	controllerRegistry := container.NewControllerRegistry(authController, feedController)
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
