package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"smartPost/models"
	"smartPost/services"
	"strconv"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{UserService: userService}
}

func (c *UserController) CreateUser(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{"error": err.Error()})
	}
	if err := c.UserService.CreateUser(&user); err != nil {

		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return ctx.JSON(http.StatusCreated, user)
}
func (c *UserController) GetUsers(ctx echo.Context) error {
	users, err := c.UserService.GetUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, users)
}
func (c *UserController) UpdateName(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	newUserName := ctx.FormValue("name")
	if newUserName == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": "UserName cannot empty"})
	}
	updateName, err := c.UserService.UpdateName(userID, ctx.FormValue("name"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"Error": "Internal server error"})
	}
	return ctx.JSON(http.StatusOK, updateName)
}
func (c *UserController) UpdateEmail(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))
	newUserEmail := ctx.FormValue("email")
	if newUserEmail == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": "Email cannot empty"})
	}
	if c.UserService.IsEmailTaken(newUserEmail) {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"Error": "Email already taken"})
	}
	updateEmail, err := c.UserService.UpdateEmail(userID, ctx.FormValue("email"))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"Error": "Internal server error"})
	}
	return ctx.JSON(http.StatusOK, updateEmail)
}
func (c *UserController) DeleteUser(ctx echo.Context) error {
	userID, _ := strconv.Atoi(ctx.Param("id"))

	if err := c.UserService.DeleteUser(userID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return ctx.NoContent(http.StatusNoContent)
}
