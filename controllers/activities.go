package controllers

import (
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ActivityHandler struct {
	DB *gorm.DB
}

func (u *ActivityHandler) GetAll(c echo.Context) error {
	activityService := services.ActivityServices{DB: u.DB, C: c}
	res, err := activityService.GetAll()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch activity data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *ActivityHandler) GetByID(c echo.Context) error {
	activityService := services.ActivityServices{DB: u.DB, C: c}
	res, err := activityService.GetByID()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{	
		"message": "Fetch activity data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *ActivityHandler) Create(c echo.Context) error {
	activityService := services.ActivityServices{DB: u.DB, C: c}

	res, err := activityService.Create()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create activity successfully",
		"data": res,
	}

	return c.JSON(http.StatusCreated, response)
}