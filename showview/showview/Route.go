package showview

import (
	"github.com/labstack/echo"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	shopmall := &Shopmall{}
	e.GET("/shopmalls", shopmall.Get)
	e.GET("/shopmalls/:id", shopmall.GetbyId)
	e.POST("/shopmalls", shopmall.Persist)
	e.PUT("/shopmalls/:id", shopmall.Put)
	e.DELETE("/shopmalls/:id", shopmall.Remove)
	return e
}
