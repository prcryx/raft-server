package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prcryx/raft-server/cmd/di/container"
	"github.com/prcryx/raft-server/internal/application/apis/auth"
	"github.com/prcryx/raft-server/internal/common/constants/routesconst"
	"github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/domain/valobj"
)

func MountAll(root chi.Router, version valobj.AppVersion, controllerRegistry *container.ControllerRegistry) {
	router := chi.NewRouter()

	// all handlers here
	router.Get(routesconst.HealthCheck, func(w http.ResponseWriter, req *http.Request) {
		utils.ResponseWithJSONData(w, http.StatusOK, struct{}{})
	})

	auth.AuthRouter(router, controllerRegistry.AuthController)

	// end of routes

	// mount
	root.Mount(version.ToString(), router)
}
