package utils

import (
	"encoding/json"
	"log"
	"net/http"

	constants "github.com/prcryx/raft-server/internal/common/constants"
	e "github.com/prcryx/raft-server/internal/common/err"
)

func ResponseWithJSONData(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add(constants.ContentType, constants.ApplicationJson)
	w.WriteHeader(code)
	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("5xx Error: %v", msg)
	}

	ResponseWithJSONData(w, code, e.ErrorResponse{
		Error: msg,
	})
}
