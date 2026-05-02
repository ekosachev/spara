package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ekosachev/spara/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	_, err := database.ConnectToDb()

	if err != nil {
		logger.Error("Could not connect to database", slog.String("error", err.Error()))
		return
	}

	globalApi := r.Group("/api")

	v1Group := globalApi.Group("/v1")

	{
		v1Group.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"Success": true})
		})
	}

	r.Run()
}
