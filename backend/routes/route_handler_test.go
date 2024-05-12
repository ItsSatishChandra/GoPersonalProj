package routes

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func TestRouteManager(t *testing.T) {

	serverTest := echo.New()
	RouteManager(serverTest)
}
