package utils

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderComponent(t templ.Component, e echo.Context) error {
	return t.Render(e.Request().Context(),e.Response().Writer);
}