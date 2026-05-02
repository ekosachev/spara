package handlers

import (
	"github.com/ekosachev/spara/internal/dto"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, dto.ApiResponse{
		Success: false,
		Error:   message,
	})
}
