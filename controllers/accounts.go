package controllers

import (
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountHandler struct {
	DB *gorm.DB
}

func (u *AccountHandler) GetAll(c echo.Context) error {
	accountService := services.AccountServices{DB: u.DB, C: c}
	res, err := accountService.GetAll()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch account data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AccountHandler) GetByID(c echo.Context) error {
	accountService := services.AccountServices{DB: u.DB, C: c}
	res, err := accountService.GetByID()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch account data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AccountHandler) Create(c echo.Context) error {
	accountService := services.AccountServices{DB: u.DB, C: c}

	res, err := accountService.Create()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create account data successfully",
		"data": res,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *AccountHandler) Update(c echo.Context) error {
	accountService := services.AccountServices{DB: u.DB, C: c}
	res, err := accountService.Update()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Update account data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *AccountHandler) Delete(c echo.Context) error {
	accountService := services.AccountServices{DB: u.DB, C: c}
	res, err := accountService.Delete()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete account data successfully",
		"data": res,
	}

	return c.JSON(http.StatusNoContent, response)
}