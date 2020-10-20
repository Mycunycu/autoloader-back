package routes

import (
	"autoposter/controllers"
	"autoposter/models"
	"net/http"
)

var userRoutes = []models.Route{
	models.Route{
		URL:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	models.Route{
		URL:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
}
