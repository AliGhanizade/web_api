package api

import (
	"fmt"

	"github.com/AliGhanizade/web_api/api/router"
	"github.com/AliGhanizade/web_api/config"
	"github.com/gin-gonic/gin"
)

func InitServer()  {
  	cfg := config.GetConfig()
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		health := v1.Group("/health")
		router.HealthRouter(health)
	}

	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))




}