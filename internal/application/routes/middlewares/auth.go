package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/prcryx/raft-server/internal/common/constants"
	e "github.com/prcryx/raft-server/internal/common/err"
	"github.com/prcryx/raft-server/internal/common/utils"
	"github.com/prcryx/raft-server/internal/infrastructure/jwt"
)

func AuthMiddleWare(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			sep := " "
			authHeader := strings.Split(r.Header.Get(constants.Authorization), sep)
			if ok := utils.ArrayLengthValidator(authHeader, 2); !ok {
				log.Println("Malformed token")

				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Malformed Token"))
			} else {
				_, err := jwt.DecodeToken(authHeader[1])
				if err != nil {
					if appErr, ok := err.(e.AppError); !ok {
						utils.ResponseWithError(w, http.StatusUnauthorized, err.Error())
					} else {
						utils.ResponseWithError(w, appErr.GetCode(), appErr.Error())
					}
					return
				}
				// ctx := context.WithValue(r.Context(), constants.Props, customClaims)
				next.ServeHTTP(w, r)
			}
		},
	)

}
