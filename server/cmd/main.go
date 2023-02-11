package main

import (
	"os"
	"snakeish/internal/handlers"
	"snakeish/internal/services/clients"
	"snakeish/pkg/core"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitCore()
	clients.Init()

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	handlers.Init(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("Listening on port: " + port)
	router.Run(":" + port)
}
