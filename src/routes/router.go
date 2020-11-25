package routes

import (
	"autoposter/controllers"
	"autoposter/middlewares"
	"github.com/gin-gonic/gin"
)

// Router - ...
type Router struct{}

// InitRouter - ...
func (r *Router) InitRouter() *gin.Engine {
	router := gin.Default()

	corsMiddleware := new(middlewares.Cors).Create()
	router.Use(corsMiddleware)

	authController := new(controllers.AuthController)
	auth := router.Group("/auth")
	{
		auth.POST("/signUp", authController.SignUp)
		auth.POST("/signIn", authController.SignIn)
	}

	return router
}
