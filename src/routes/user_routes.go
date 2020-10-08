package routes

import (
	"net/http"
	"autoposter/controllers"
	"autoposter/models"
)

var userRoutes = []models.Route {
	models.Route {
		URL: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	models.Route {
		URL: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
}