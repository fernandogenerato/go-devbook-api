package router

import (
	"github.com/gorilla/mux"
	"go-devbook-api/src/router/routes"
)

func Build() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
