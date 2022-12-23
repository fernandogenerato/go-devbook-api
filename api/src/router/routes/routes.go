package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"go-devbook-api/src/middlewares"
)

type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	for _, route := range routes {

		if route.Authentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.DoAuth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
