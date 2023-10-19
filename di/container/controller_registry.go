package container

import "github.com/prcryx/raft-server/internal/application/apis/auth"

type ControllerRegistry struct {
	AuthController *auth.AuthController
}

func NewControllerRegistry(authController *auth.AuthController) *ControllerRegistry {
	return &ControllerRegistry{
		AuthController: authController,
	}
}
