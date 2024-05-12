package backend

import (
	"github.com/labstack/echo/v4"

	"website/backend/routes"
)

func Server() error {
	server := echo.New()

	routes.RouteManager(server)

	err := server.Start(":8000")

	if err != nil {
		return err
	}
	return nil
}
