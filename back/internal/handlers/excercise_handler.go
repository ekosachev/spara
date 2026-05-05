package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/services"
	"github.com/gin-gonic/gin"
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
