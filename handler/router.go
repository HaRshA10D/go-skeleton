package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/user-name/skeleton-name/service"
	"log"
	"net/http"
	"strings"
)

func NewRouter(useCases service.UseCases) http.Handler {

	router := mux.NewRouter()
	rootRouter := router.PathPrefix("/").Subrouter()

	initPingRoutes(rootRouter)

	if err := printRoutes(router); err != nil {
		log.Printf("error in printing routes: %s\n", err)
	}
	return router
}

func printRoutes(router *mux.Router) error {
	log.Println("Printing all http routes from router")
	return router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(strings.Repeat("| ", len(ancestors)), t)
		return nil
	})
}
