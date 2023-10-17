package routers

import (
	"apiwrap/handlers"
	"apiwrap/middlewares"
	"github.com/gin-gonic/gin"
)

func R(e *gin.Engine) {
	e.Use(middlewares.CORS())

	e.GET("/", handlers.HealthCheck)
	e.NoRoute(handlers.Wrap)
}
