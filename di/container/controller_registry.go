package container

import (
	"github.com/prcryx/raft-server/internal/application/apis/auth"
	feed "github.com/prcryx/raft-server/internal/application/apis/feeds"
)

type ControllerRegistry struct {
	AuthController *auth.AuthController
	FeedController *feed.FeedController
}

func NewControllerRegistry(authController *auth.AuthController, feedController *feed.FeedController) *ControllerRegistry {
	return &ControllerRegistry{
		AuthController: authController,
		FeedController: feedController,
	}
}
