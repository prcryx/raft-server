package wire

import (
	"github.com/google/wire"
	ri "github.com/prcryx/raft-server/internal/data/repository_impl"
	"github.com/prcryx/raft-server/internal/domain/repositories"
)

var (
	AuthRepositorySet = wire.NewSet(
		ri.NewAuthRepositoryImpl,
		wire.Bind(
			new(repositories.AuthRepository),
			new(*ri.AuthRepositoryImpl),
		),
	)
)

var (
	RepositorySet = wire.NewSet(
		AuthRepositorySet,
	)
)
