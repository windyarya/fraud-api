package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"bitbucket.org/windyarya/backend-final/utils/token"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServices struct {	
	DB *gorm.DB
	C echo.Context
}

func (u *UserServices) GetUsers() ([]models.User, error) {
	var users []models.User

	err := u.DB.Find(&users)
	println(err)
	if err.Error != nil {
		println(err)
		return users, err.Error
	}
	
	return users, nil
}

func (u *UserServices) GetUser() (models.User, error) {
	var user models.User
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return user, err
	}
	
	err2 := u.DB.Find(&user, ids)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return (models.User{}), err2.Error
	}
	
	return user, nil
}

func (u *UserServices) Register() (models.User, error) {
	var user models.User
	err := u.C.Bind(&user)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return (models.User{}), err
	}

	pass, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(pass)
	
	err2 := u.DB.Create(&user)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return user, err2.Error
	}
	
	return user, nil
}

func (u *UserServices) Login() (string, error) {
	var user models.User

	err := u.C.Bind(&user)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return "", err
	}

	reqPass := user.Password

	err2 := u.DB.Where("email = ?", user.Email).First(&user)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return "", err2.Error
	}

	_ = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPass))

	token, _ := token.ClaimToken(user.ID, user.UserGroupID, user.WorkUnitID)

	return token, err2.Error
}

func (u *UserServices) UpdateUser() (models.User, error) {
	var user models.User
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return user, err
	}

	err2 := u.C.Bind(&user)
	if err2 != nil {
		u.C.Logger().Error(err2.Error)
		return user, err2
	}

	err3 := u.DB.Model(&user).Where("id = ?", ids).Updates(&user)
	if err3.Error != nil {
		u.C.Logger().Error(err3.Error)
		return user, err3.Error
	}

	var updatedUser models.User
	err4 := u.DB.Find(&updatedUser, ids)
	if err4.Error != nil {
		u.C.Logger().Error(err4.Error)
		return user, err4.Error
	}

	return updatedUser, nil
}

func (u *UserServices) DeleteUser() (models.User, error) {
	var user models.User
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return user, err
	}

	err2 := u.DB.Delete(&user, ids)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return user, err2.Error
	}

	return user, nil
}