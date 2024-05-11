package hello

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Hello(context echo.Context) error {
	return context.String(http.StatusOK, "hello world")
}
