package routes

import (
	"github.com/Igorsant/go-api/internal/config"
	"github.com/Igorsant/go-api/internal/hello"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	config := config.NewConfig()
	helloController := hello.NewHelloController(config)
	helloController.Handle(r)
}
