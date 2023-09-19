package router

import (
	"github.com/fulunageli/gofarm/router/http"
	"github.com/labstack/echo/v4"
)

func LoadAllRouters(app *echo.Echo) {
	http.IndexRouter(app)
}
