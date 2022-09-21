package router

import (
	"net/http"

	"github.com/julianstephens/stripe-shopping-app-tutorial/server/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	api := e.Group("/api")
	{
		api.GET("/ping", func(c echo.Context) error {
			return c.JSON(http.StatusOK, &utils.HttpResp{Message: "Ping received", Data: nil})
		})
	}

	return e
}
