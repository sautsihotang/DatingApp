package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sautsihotang/DatingApp/controllers"
)

func Route(e *echo.Echo) {

	e.POST("api/auth/register", controllers.Register)
	e.POST("api/auth/login", controllers.Login)
}
