package routes

import (
	"autoposter/models"
	"github.com/gorilla/mux"
)

// NewRouter - creating new Router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	return SetupRoute(r)
}

// LoadAllRoutes - ...
func LoadAllRoutes() []models.Route {
	routes := userRoutes
	return routes
}

// SetupRoute - creating new route with settings
func SetupRoute(r *mux.Router) *mux.Router {
	for _, route := range LoadAllRoutes() {
		r.HandleFunc(route.URL, route.Handler).Methods(route.Method)
	}
	return r
}
