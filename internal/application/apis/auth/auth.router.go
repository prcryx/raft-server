package auth

import (
	"github.com/go-chi/chi/v5"
	"github.com/prcryx/raft-server/internal/common/constants/routesconst"
)

func AuthRouter(router *chi.Mux, ac *AuthController) {
	router.Post(routesconst.SignUp, ac.SignUpWithEmailAndPassword)
}
