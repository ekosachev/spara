package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/services"
	"github.com/ekosachev/spara/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
	logger      *slog.Logger
}

func NewAuthHandler(userService *services.UserService, logger *slog.Logger) AuthHandler {
	return AuthHandler{userService, logger}
}

func (h *AuthHandler) RegisterRoutes(group *gin.RouterGroup, prefix string) {
	routerGroup := group.Group(prefix)

	{
		routerGroup.POST("/login", h.login)
	}
}

func (h *AuthHandler) login(c *gin.Context) {
	var req dto.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.userService.GetByEmail(c, req.Email)

	if err != nil {
		sendError(c, http.StatusNotFound, "Incorrect email or password")
		return
	}

	if utils.CheckPassword(req.Password, user.PasswordHash) != nil {
		sendError(c, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{Success: true})
}
