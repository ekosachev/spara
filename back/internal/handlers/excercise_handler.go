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
	localGroup := group.Group(prefix)

	{
		localGroup.GET("/", h.getAll)
	}
}

func (h *ExcerciseHandler) getAll(c *gin.Context) {
	excercises, err := h.service.GetAll(c)

	if err != nil {
		sendError(c, http.StatusInternalServerError, "Internal server error")
		h.logger.Error("Failed to get excercises", slog.String("error", err.Error()))
	}

	response := make([]dto.ExcerciseResponse, len(excercises))
	for i, ex := range excercises {
		response[i] = dto.ExcerciseResponse{
			ID:          ex.ID,
			Name:        ex.Name,
			Description: ex.Description,
		}
	}

	c.JSON(http.StatusOK, dto.ApiResponse{Success: true, Data: response})
}
