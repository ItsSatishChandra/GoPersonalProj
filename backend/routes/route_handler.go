package routes

import (
	"website/backend/apps/hello"

	"github.com/labstack/echo/v4"
)

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

func RouteManager(server *echo.Echo) {
	server.Match([]string{POST, GET}, RouteNameResolver("hello.Hello"), hello.Hello)
}
