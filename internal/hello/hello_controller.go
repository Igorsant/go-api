package hello

import (
	"github.com/Igorsant/go-api/internal/config"
	"github.com/gin-gonic/gin"
)

type HelloController struct {
	config config.Config
}

func NewHelloController(config config.Config) HelloController {
	return HelloController{
		config: config,
	}
}

func (h *HelloController) Handle(r *gin.Engine) {
	r.GET(h.config.GetHealthRoute, func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "healthy",
		})
	})
}
