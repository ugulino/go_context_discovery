package main

import (
	"github.com/gin-gonic/gin"

	"goContextDiscovery/cmd/internal/core/services"
	"goContextDiscovery/cmd/internal/handlers"
	"goContextDiscovery/cmd/internal/repositories"
)

func main() {
	repo := repositories.NewMemKVS()
	service := services.NewInstance(repo)
	handler := handlers.NewHttpHandler(service)

	router := gin.New()
	router.GET("/person/:id", handler.Find)
	router.POST("/persons", handler.SignUp)

	router.Run(":8080")
}
