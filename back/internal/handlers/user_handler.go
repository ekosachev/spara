package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ekosachev/spara/internal/dto"
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

func (h *UserHandler) RegisterRoutes(group *gin.RouterGroup, prefix string) {
	routerGroup := group.Group(prefix)

	{
		routerGroup.POST("/", h.create)
	}
}

func (h *UserHandler) create(c *gin.Context) {
	var req dto.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.Create(c, dto.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: req.Password, // implement hashing later
	}); err != nil {
		h.logger.Error("Failed to create user", slog.String("error", err.Error()))
		sendError(c, http.StatusInternalServerError, "Internal server error")
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Success: true,
		Data:    "User created",
	})
}
