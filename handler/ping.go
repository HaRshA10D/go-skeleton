package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

func initPingRoutes(router *mux.Router) {
	router.HandleFunc("/ping", handlePing()).Methods(http.MethodGet)
}

func handlePing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`pong`))
	}
}
