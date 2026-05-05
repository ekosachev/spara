package handlers

import (
	"github.com/ekosachev/spara/internal/services"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type ExcerciseHandler struct {
	service *services.ExcerciseService
	logger  *slog.Logger
}

func NewExcerciseHandler(service *services.ExcerciseService, logger *slog.Logger) ExcerciseHandler {
	return ExcerciseHandler{service, logger}
}

func (h *ExcerciseHandler) RegisterRoutes(group *gin.RouterGroup, prefix string) {
	_ = group.Group(prefix)

	{
	}
}
