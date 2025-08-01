package httpadapter

import (
	"api/internal/application"
	"api/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// CreateEvent creates a new event
// @summary Create an event// 
// @Description Create a new event
// @Produce json
// @Router /events [post]
// @Tags Events
// @Accept json
// @Param user body domain.Event true "Event data"
// @Success 201 {object} domain.Event "'Event' created" 
func (h *Handler) CreateEvent(c *gin.Context) {
	var event domain.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	event.ID = uuid.New().String()
	if err := h.EventService.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, event)
}
 
// CreateUser creates a new user
// @summary Create a user
// @Description Create a new user
// @Produce json
// @Router /users [post]
// @Tags Users
// @Accept json
// @Param user body domain.User true "User data"
// @Success 201 {object} domain.User "User created" 
func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = uuid.New().String()
	if err := h.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
