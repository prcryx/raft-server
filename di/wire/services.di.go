package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/infrastructure/jwt"
)

var (
	JwtStrategy = wire.NewSet(
		jwt.NewJwtStrategy,
		wire.Bind(
			new(jwt.IJwtStrategy),
			new(*jwt.JwtStrategy),
		),
	)
)

var (
	OtherServices = wire.NewSet(
		JwtStrategy,
	)
)
