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
	e.PUT("/users/:id/editRole", userController.UpdateRole)
	e.PUT("/users/:id/createUserGroup", userController.CreateUserGroup)
	e.DELETE("users/deleteUserGroup/", userController.DeleteUserGroup)
	e.GET("/users/getUserGroup", userController.GetUserGroups)
	e.GET("/groups", userController.GetGroups)
	e.DELETE("/users/:id", userController.DeleteUser)
}
