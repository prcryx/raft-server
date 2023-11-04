package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/domain/usecases"
)

var (
	AuthUseCaseSet = wire.NewSet(
		usecases.NewAuthUseCase,
		wire.Bind(
			new(usecases.IAuthUseCase),
			new(*usecases.AuthUseCase),
		),
	)
	FeedUseCaseSet = wire.NewSet(
		usecases.NewFeedUseCase,
		wire.Bind(
			new(usecases.IFeedUseCase),
			new(*usecases.FeedUseCase),
		),
	)
)

var (
	UseCaseSet = wire.NewSet(
		AuthUseCaseSet,
		FeedUseCaseSet,
	)
)
