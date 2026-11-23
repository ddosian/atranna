package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ddosian/atranna/api/docs"
	"github.com/ddosian/atranna/api/internal/db"
	"github.com/ddosian/atranna/api/internal/server/routers"
)

// @title           Atranna API
// @version         1.0
// @description     API backend for Atranna
// @BasePath        /api/v1
func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "/data/atranna.db"
	}
	conn, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("open database: %v", err)
	}
	defer conn.Close()

	router := gin.Default()

	v1Router := router.Group("/api/v1")

	v1Router.GET("/health", routers.GetHealth(conn))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("0.0.0.0:8080")
}
