package app

import (
	"github.com/prcryx/raft-server/config"
	"github.com/prcryx/raft-server/di/container"
	"github.com/prcryx/raft-server/internal/domain/valobj"
)

type App struct {
	ControllerRegistry *container.ControllerRegistry
	ServicesRegistry   *container.ServicesRegistry
	EnvConfig          *config.EnvConfig
	Version            valobj.AppVersion
}

func NewApp(controllerRegistry *container.ControllerRegistry, servicesRegistry *container.ServicesRegistry, envConfig *config.EnvConfig, version string) *App {
	return &App{
		ControllerRegistry: controllerRegistry,
		ServicesRegistry:   servicesRegistry,
		EnvConfig:          envConfig,
		Version:            valobj.GetAppVersion(version),
	}
}
