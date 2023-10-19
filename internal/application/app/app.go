package app

import (
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	"github.com/prcryx/raft-server/internal/domain/valobj"
)

type App struct {
	ControllerRegistry *container.ControllerRegistry

	EnvConfig *config.EnvConfig
	Version   valobj.AppVersion
}

func NewApp(controllerRegistry *container.ControllerRegistry, envConfig *config.EnvConfig, version string) *App {
	return &App{
		ControllerRegistry: controllerRegistry,

		EnvConfig: envConfig,

		Version: valobj.GetAppVersion(version),
	}
}
