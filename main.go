package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"smartPost/controllers"
	"smartPost/repositories"
	"smartPost/routes"
	"smartPost/services"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	connectString := "host=localhost port=5432 dbname=testing user=postgres password=hwangthai2001 sslmode=disable TimeZone=Asia/Ho_Chi_Minh"
	db, err := gorm.Open("postgres", connectString)
	if err != nil {
		fmt.Println("Failed to connect database", err)
		return
	}
	defer db.Close()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	routes.SetupRoutes(e, userController)
	e.Logger.Fatal(e.Start(":8888"))
}
