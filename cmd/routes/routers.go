package routes

import (
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/prcryx/raft-server/internal/common/constants"
	"github.com/prcryx/raft-server/internal/common/utils"
)

func Root(firebaseInstance *firebase.App) chi.Router {
	rootRoute := chi.NewRouter()
	rootRoute.Use(cors.Handler(
		cors.Options{
			AllowedOrigins:   constants.AllowedOrigins(),
			AllowedHeaders:   constants.AllowedHeaders(),
			AllowedMethods:   constants.AllowedMethods(),
			AllowCredentials: constants.AllowCredentials,
			MaxAge:           constants.MaxAge,
			ExposedHeaders:   constants.ExposedHeaders(),
		},
	))

	return rootRoute
}

func MountAll(root chi.Router, mountPath string) {
	router := chi.NewRouter()

	// all handlers here
	router.Get(HealthCheck, func(w http.ResponseWriter, req *http.Request) {
		utils.ResponseWithJSONData(w, http.StatusOK, struct{}{})
	})
	//auth router


	root.Mount(mountPath, router)
}
