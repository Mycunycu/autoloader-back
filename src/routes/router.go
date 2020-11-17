package routes

import (
	"autoposter/controllers"
	"github.com/gin-gonic/gin"
)

// Router - ...
type Router struct{}

// InitRouter - ...
func (r *Router) InitRouter() *gin.Engine {
	router := gin.New()

	authCtrl := new(controllers.AuthController)
	auth := router.Group("/auth")
	{
		auth.POST("/signUp", authCtrl.SignUp)
		auth.POST("/signIn", authCtrl.SignIn)
	}

	return router
}
