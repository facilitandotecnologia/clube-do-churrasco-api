package httpadapter

import (
	"api/internal/application"
	"api/internal/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	EventService *application.EventService
	UserService  *application.UserService
}

func NewHandler(eventService *application.EventService, userService *application.UserService) *Handler {
	return &Handler{
		EventService: eventService,
		UserService:  userService,
	}
}

// Event Handlers

func (h *Handler) CreateEvent(c *gin.Context) {
	var event domain.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = time.Now().Format("20060102150405")
	if err := h.EventService.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, event)
}

// User Handlers

func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = time.Now().Format("20060102150405")
	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
