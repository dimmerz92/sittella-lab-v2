package pages

import (
	"github.com/labstack/echo/v4"
	"sittellalab.com.au/views"
)

func Home(ctx echo.Context) error {
	return Render(ctx, views.Home())
}
