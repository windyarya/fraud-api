package main

import (
	"bitbucket.org/windyarya/backend-final/database"
	"bitbucket.org/windyarya/backend-final/database/migrations"
	"bitbucket.org/windyarya/backend-final/routes"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	migration.Migration(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5173", "http://localhost:5173", "http://ec2-13-215-179-160.ap-southeast-1.compute.amazonaws.com", "http://ec2-13-215-179-160.ap-southeast-1.compute.amazonaws.com:80", "http://ec2-13-215-179-160.ap-southeast-1.compute.amazonaws.com:8080", "http://ec2-13-215-179-160.ap-southeast-1.compute.amazonaws.com:443"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	routes.Router(e, db)

	e.Logger.Fatal(e.Start(":8080"))
}