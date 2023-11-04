package routes

import (
	// "net/http"
	"github.com/go-chi/chi/v5"
	// "github.com/prcryx/raft-server/cmd/di/container"
	"github.com/prcryx/raft-server/internal/application/app"
	"github.com/prcryx/raft-server/internal/application/routes/middlewares"

	// "github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/domain/types"
	// "github.com/prcryx/raft-server/internal/domain/valobj"
	// "github.com/prcryx/raft-server/internal/apis/auth"
)

func SetupRoutes(app *app.App, server *types.Server) {
	root := chi.NewRouter()
	root.Use(middlewares.Cors())
	root.Use(middlewares.Logger)
	MountAll(root, app.Version, app.ControllerRegistry,)
	// MountAll(root, app.Version, app.ControllerRegistry)
	server.Router = root
}
