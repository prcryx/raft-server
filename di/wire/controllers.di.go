package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/application/apis/auth"
)

var (
	AuthControllerSet = wire.NewSet(
		auth.NewAuthController,
		wire.Bind(
			new(auth.IAuthController),
			new(*auth.AuthController),
		),
	)
)

var (
	ControllerSet = wire.NewSet(
		AuthControllerSet,
	)
)
