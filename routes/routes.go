package routes

import (
	"github.com/labstack/echo/v4"
	"smartPost/controllers"
)

func SetupRoutes(e *echo.Echo, userController *controllers.UserController) {
	e.GET("/users", userController.GetUsers)
	e.POST("/users", userController.CreateUser)
	e.PUT("/users/:id/editName", userController.UpdateName)
	e.PUT("/users/:id/editEmail", userController.UpdateEmail)
	e.DELETE("users/:id", userController.DeleteUser)
}
