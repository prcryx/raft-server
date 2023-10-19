package wire

import (
	"github.com/google/wire"
	"github.com/prcryx/raft-server/internal/data/datasoruces"
)

var (
	AuthDataSourceTest = wire.NewSet(
		datasoruces.NewAuthDataSource,
		wire.Bind(
			new(datasoruces.IAuthDataSource),
			new(*datasoruces.AuthDataSource),
		),
	)
)

var (
	DataSourceSet = wire.NewSet(
		AuthDataSourceTest,
	)
)
