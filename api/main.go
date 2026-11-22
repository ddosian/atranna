package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ddosian/atranna/api/docs"
	"github.com/ddosian/atranna/api/internal/server/routers"
)

// @title           Atranna API
// @version         1.0
// @description     API backend for Atranna
// @BasePath        /api/v1
func main() {
	router := gin.Default()

	v1_router := router.Group("/api/v1")

	v1_router.GET("/health", routers.GetHealth)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:8080")
}
