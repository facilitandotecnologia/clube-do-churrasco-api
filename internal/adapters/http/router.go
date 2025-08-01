package httpadapter

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *Handler) {
	r.POST("/events", handler.CreateEvent)
	r.POST("/users", handler.CreateUser)
	// Add more routes as needed
}
