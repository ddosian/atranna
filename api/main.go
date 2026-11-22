package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ddosian/atranna/api/internal/server/routers"
)

func main() {
	router := gin.Default()

	v1_router := router.Group("/api/v1")

	v1_router.GET("/health", routers.GetHealth)

	router.Run("0.0.0.0:8080")
}
