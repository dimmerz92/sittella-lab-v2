package pages

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(ctx echo.Context, t templ.Component) error {
	return t.Render(ctx.Request().Context(), ctx.Response())
}
