package main

import (
	"github.com/fulunageli/gofarm/constant"
	"github.com/fulunageli/gofarm/router"
	"github.com/fulunageli/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	//Load static files
	constant.LoadStatic(app)

	//Load template folder
	app.Renderer = constant.LoadTemplate()

	// router
	router.LoadAllRouters(app)

	//server
	server.SetServer(app)
}
