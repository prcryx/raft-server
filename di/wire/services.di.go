package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/infrastructure/jwt"
)

var (
	JwtStrategySet = wire.NewSet(
		jwt.NewJwtStrategy,
		wire.Bind(
			new(jwt.IJwtStrategy),
			new(*jwt.JwtStrategy),
		),
	)
)

var (
	OtherServicesSet = wire.NewSet(
		JwtStrategySet,
	)
)
