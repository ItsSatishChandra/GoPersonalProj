package backend

import (
	"github.com/labstack/echo/v4"

	"website/backend/routes"
)

func Server() {
	server := echo.New()

	routes.RouteManager(server)

	server.Logger.Fatal(server.Start(":8000"))
}
