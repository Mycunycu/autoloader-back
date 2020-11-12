package routes

import (
	"github.com/gin-gonic/gin"
)

// Router - ...
type Router struct {}

// InitRouter - ...
func (r *Router) InitRouter() *gin.Engine {
	router := gin.New()
	
	auth := router.Group("/auth", ) 
	{
		auth.POST("/sign-up",)
		auth.POST("/sign-in",)
	}

	return router
}
