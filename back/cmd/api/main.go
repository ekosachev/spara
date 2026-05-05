package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ekosachev/spara/internal/database"
	"github.com/ekosachev/spara/internal/handlers"
	"github.com/ekosachev/spara/internal/repositories"
	"github.com/ekosachev/spara/internal/services"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.ConnectToDb()

	if err != nil {
		logger.Error("Could not connect to database", slog.String("error", err.Error()))
		return
	}

	excerciseRepo := repositories.NewExcerciseRepository(db)

	if err := database.Seed(&excerciseRepo); err != nil {
		logger.Error("Failed to seed database", slog.String("error", err.Error()))
	}

	globalApi := r.Group("/api")

	v1Group := globalApi.Group("/v1")

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(&userService, logger)

	authHandler := handlers.NewAuthHandler(&userService, logger)

	{
		v1Group.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"Success": true})
		})

		userHandler.RegisterRoutes(v1Group, "/user")
		authHandler.RegisterRoutes(v1Group, "/auth")
	}

	r.Run()
}
