package handlers

import (
	"log/slog"
	"net/http"

	"github.com/ekosachev/spara/internal/dto"
	"github.com/ekosachev/spara/internal/middleware"
	"github.com/ekosachev/spara/internal/services"
	"github.com/gin-gonic/gin"
)

type TrainingPlanHandler struct {
	service *services.TrainingPlanServcie
	logger  *slog.Logger
}

func NewTrainingPlanHandler(service *services.TrainingPlanServcie, logger *slog.Logger) TrainingPlanHandler {
	return TrainingPlanHandler{service, logger}
}

func (h *TrainingPlanHandler) RegisterRoutes(group *gin.RouterGroup, prefix string) {
	localGroup := group.Group(prefix)

	{
		protectedGroup := localGroup.Group("/").Use(middleware.AuthMiddleware())

		protectedGroup.POST("/", h.create)
	}
}

func (h *TrainingPlanHandler) create(c *gin.Context) {
	var req dto.CreateTrainingPlanRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.GetUint("userID")

	trainingPlan, err := h.service.Create(c, dto.TrainingPlan{Name: req.Name, UserID: userID})

	if err != nil {
		sendError(c, http.StatusInternalServerError, "Internal server error")
		h.logger.Error("Failed to create training plan", slog.String("error", err.Error()))
		return
	}

	c.JSON(http.StatusCreated, dto.ApiResponse{Success: true, Data: dto.TrainingPlanResponse{
		Name:   trainingPlan.Name,
		UserID: trainingPlan.UserID,
	}})
}
