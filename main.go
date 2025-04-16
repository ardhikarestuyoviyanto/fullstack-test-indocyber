package main

import (
	"auth_test/config"
	"auth_test/internal/db"
	"auth_test/internal/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.InitConfiguration()
	e:= echo.New()
	routes.RegisterAuthAPI(e, db.InitDb(), config)
	e.Logger.Fatal(e.Start("8000"))
}