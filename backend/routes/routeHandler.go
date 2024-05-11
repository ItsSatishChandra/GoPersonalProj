package routes

import (
	"website/backend/apps/hello"

	"github.com/labstack/echo/v4"
)

func RouteManager(server *echo.Echo) {
	server.GET("/", hello.Hello)
}
