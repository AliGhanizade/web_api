package router

import (
	"github.com/AliGhanizade/web_api/api/handlers"
	"github.com/gin-gonic/gin"
)

func  HealthRouter(r *gin.RouterGroup) {
	handlers := handlers.NewHealthHandler()
	
		r.GET("/", handlers.Health)
		r.POST("/:id", handlers.HealthByID)

	
}