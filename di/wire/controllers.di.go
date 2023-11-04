package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/application/apis/auth"
	feed "github.com/prcryx/raft-server/internal/application/apis/feeds"
)

var (
	AuthControllerSet = wire.NewSet(
		auth.NewAuthController,
		wire.Bind(
			new(auth.IAuthController),
			new(*auth.AuthController),
		),
	)

	FeedControllerSet = wire.NewSet(
		feed.NewFeedController,
		wire.Bind(
			new(feed.IFeedController),
			new(*feed.FeedController),
		),
	)
)

var (
	ControllerSet = wire.NewSet(
		AuthControllerSet,
		FeedControllerSet,
	)
)
