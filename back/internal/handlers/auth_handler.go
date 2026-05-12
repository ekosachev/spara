package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/middleware"
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

		protectedGroup := routerGroup.Group("").Use(middleware.AuthMiddleware())
		{
			protectedGroup.GET("/me", h.me)
		}
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
		sendError(c, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	if utils.CheckPassword(req.Password, user.PasswordHash) != nil {
		sendError(c, http.StatusUnauthorized, "Incorrect email or password")
		return
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		sendError(c, http.StatusInternalServerError, "Internal server error")
		h.logger.Error("Failed to generate JWT token", slog.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Success: true,
		Data:    dto.LoginResponse{Token: token},
	})
}

func (h *AuthHandler) me(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		sendError(c, http.StatusUnauthorized, "User ID not found")
		return
	}

	user, err := h.userService.GetByID(c, uint(userID.(float64)))

	if err != nil {
		sendError(c, http.StatusNotFound, "User does not exist")
		return
	}

	c.JSON(http.StatusOK, dto.ApiResponse{
		Success: true,
		Data: dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	})
}
