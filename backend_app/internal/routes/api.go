package routes

import (
	"auth_test/config"
	"auth_test/internal/handler"
	"auth_test/internal/repository"
	"auth_test/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAuthAPI(e *echo.Echo, db *gorm.DB, conf config.Config) {
	repo := repository.NewAuthRepository(db)
	svc := service.NewServiceAuth(repo, conf)
	cont := handler.AuthEchoHandler{
		SvcAuth: svc,
	}

	e.POST("/login", cont.LoginHandler)
}