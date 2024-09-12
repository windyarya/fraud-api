package controllers

import (
	"bitbucket.org/windyarya/backend-final/models"
	"bitbucket.org/windyarya/backend-final/services"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func (u *UserHandler) GetUsers(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	res, err := userService.GetUsers()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}

		return c.JSON(http.StatusInternalServerError, data)
	}

	var userResponses []models.UserResponse
	for _, user := range res {
        userResponses = append(userResponses, models.UserResponse{
            ID:        user.ID,
            CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
            Name:      user.Name,
            Email:     user.Email,
			UserGroupID:   user.UserGroupID,
			WorkUnitID:    user.WorkUnitID,
        })
    }

	response := map[string]interface{}{
		"message": "Fetch user data successfully",
		"data": userResponses,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserHandler) GetUser(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	res, err := userService.GetUser()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	} else if res.ID == 0 {
		data := map[string]interface{}{
			"message": "User not found",
		}
		return c.JSON(http.StatusNotFound, data)
	}

	userResponse := models.UserResponse{
		ID:        res.ID,
		CreatedAt: res.CreatedAt.Format(time.RFC3339),
		UpdatedAt: res.UpdatedAt.Format(time.RFC3339),
		Name:      res.Name,
		Email:     res.Email,
		UserGroupID:   res.UserGroupID,
		WorkUnitID:    res.WorkUnitID,
	}

	response := map[string]interface{}{
		"message": "Fetch user data successfully",
		"data": userResponse,
	}

	return c.JSON(http.StatusOK, response)
}

func (u *UserHandler) Register(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	res, err := userService.Register()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	userResponse := models.UserResponse{
		ID:        res.ID,
		CreatedAt: res.CreatedAt.Format(time.RFC3339),
		UpdatedAt: res.UpdatedAt.Format(time.RFC3339),
		Name:      res.Name,
		Email:     res.Email,
		UserGroupID:   res.UserGroupID,
		WorkUnitID:    res.WorkUnitID,
	}

	response := map[string]interface{}{
		"message": "Register success",
		"data": userResponse,
	}
	return c.JSON(http.StatusCreated, response)
}

func (u *UserHandler) Login(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	res, err := userService.Login()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Login success",
		"data": res,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *UserHandler) Update(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	res, err := userService.UpdateUser()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	userResponse := models.UserResponse{
		ID:        res.ID,
		CreatedAt: res.CreatedAt.Format(time.RFC3339),
		UpdatedAt: res.UpdatedAt.Format(time.RFC3339),
		Name:      res.Name,
		Email:     res.Email,
		UserGroupID:   res.UserGroupID,
		WorkUnitID:    res.WorkUnitID,
	}

	response := map[string]interface{}{
		"message": "Register success",
		"data": userResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func (u *UserHandler) Delete(c echo.Context) error {
	userService := services.UserServices{DB: u.DB, C: c}
	_, err := userService.DeleteUser()
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Delete success",
	}
	return c.JSON(http.StatusNoContent, response)
}