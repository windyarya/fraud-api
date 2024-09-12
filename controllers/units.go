package controllers

import (
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UnitHandler struct {
	DB *gorm.DB
}

func (u *UnitHandler) GetAll(c echo.Context) error {
	unitService := services.UnitServices{DB: u.DB, C: c}
	res, err := unitService.GetAll()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch unit data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UnitHandler) GetByID(c echo.Context) error {
	unitService := services.UnitServices{DB: u.DB, C: c}
	res, err := unitService.GetByID()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch unit data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UnitHandler) Create(c echo.Context) error {
	unitService := services.UnitServices{DB: u.DB, C: c}
	res, err := unitService.Create()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create unit data successfully",
		"data": res,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *UnitHandler) Update(c echo.Context) error {
	unitService := services.UnitServices{DB: u.DB, C: c}
	res, err := unitService.Update()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Update unit data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UnitHandler) Delete(c echo.Context) error {
	unitService := services.UnitServices{DB: u.DB, C: c}
	_, err := unitService.Delete()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete unit data successfully",
	}

	return c.JSON(http.StatusNoContent, response)
}