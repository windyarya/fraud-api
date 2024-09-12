package controllers

import (
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GroupHandler struct {
	DB *gorm.DB
}

func (u *GroupHandler) GetAll(c echo.Context) error {
	groupService := services.GroupServices{DB: u.DB, C: c}
	res, err := groupService.GetAll()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch group data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *GroupHandler) GetByID(c echo.Context) error {
	groupService := services.GroupServices{DB: u.DB, C: c}
	res, err := groupService.GetByID()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Fetch group data successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *GroupHandler) Create(c echo.Context) error {
	groupService := services.GroupServices{DB: u.DB, C: c}
	res, err := groupService.Create()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Create group successfully",
		"data": res,
	}

	return c.JSON(http.StatusCreated, response)
}

func (u *GroupHandler) Update(c echo.Context) error {
	groupService := services.GroupServices{DB: u.DB, C: c}
	res, err := groupService.Update()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Update group successfully",
		"data": res,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *GroupHandler) Delete(c echo.Context) error {
	groupService := services.GroupServices{DB: u.DB, C: c}
	_, err := groupService.Delete()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete group successfully",
	}

	return c.JSON(http.StatusNoContent, response)
}