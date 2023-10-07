package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/prcryx/raft-server/common/utils"
)

func Root(middleware http.Handler) chi.Router {
	rootRoute := chi.NewRouter()
	rootRoute.Use()

	return rootRoute
}

func MountAll(root chi.Router, mountPath string) {
	router := chi.NewRouter()

	// all handlers here
	router.Get(HealthCheck, func (w http.ResponseWriter, req *http.Request) {
        utils.ResponseWithJSONData(w, http.StatusOK, struct{}{})
    })
    // app.Get("/health", func(c *fiber.Ctx) error {
	// 	return c.SendString("Healthy!")
	// })

	root.Mount(mountPath, router)
}
