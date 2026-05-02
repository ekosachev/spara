package handlers

import (
	"log/slog"

	"github.com/ekosachev/spara/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
	logger  *slog.Logger
}

func NewUserHandler(service services.UserService, logger *slog.Logger) UserHandler {
	return UserHandler{service, logger}
}

func (h *UserHandler) RegisterRoutes(group gin.RouterGroup, prefix string) {
	routerGroup := group.Group(prefix)

	{
		// register routes here
	}
}

func (h *UserHandler) Create(c *gin.Context)
