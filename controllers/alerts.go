package controllers

import (
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AlertHandler struct {
	DB *gorm.DB
}

func (u *AlertHandler) GetAll(c echo.Context) error {
	alertService := services.AlertServices{DB: u.DB, C: c}
	res, err := alertService.GetAll()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch alert data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AlertHandler) GetByID(c echo.Context) error {
	alertService := services.AlertServices{DB: u.DB, C: c}
	res, err := alertService.GetByID()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch alert data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AlertHandler) Create(c echo.Context) error {
	alertService := services.AlertServices{DB: u.DB, C: c}

	res, err := alertService.Create()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create alert successfully",
		"data": res,
	}

	return c.JSON(http.StatusNoContent, response)
}

func (u *AlertHandler) Update(c echo.Context) error {
	alertService := services.AlertServices{DB: u.DB, C: c}
	res, err := alertService.Update()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Update alert successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AlertHandler) Delete(c echo.Context) error {
	alertService := services.AlertServices{DB: u.DB, C: c}
	res, err := alertService.Delete()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete alert successfully",
		"data": res,
	}

	return c.JSON(http.StatusNoContent, response)
}