package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UnitServices struct {	
	DB *gorm.DB
	C echo.Context
}

func (u *UnitServices) GetAll() ([]models.WorkUnit, error) {
	var units []models.WorkUnit

	err := u.DB.Find(&units)
	println(err)
	if err.Error != nil {
		println(err)
		return units, err.Error
	}
	
	return units, nil
}

func (u *UnitServices) GetByID() (models.WorkUnit, error) {
	var unit models.WorkUnit
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return unit, err
	}
	
	err2 := u.DB.Find(&unit, ids)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return unit, err2.Error
	}
	
	return unit, nil
}

func (u *UnitServices) Create() (models.WorkUnit, error) {
	var unit models.WorkUnit
	err := u.C.Bind(&unit)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return unit, err
	}
	
	err2 := u.DB.Create(&unit)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return unit, err2.Error
	}
	
	return unit, nil
}

func (u *UnitServices) Update() (models.WorkUnit, error) {
	var unit models.WorkUnit
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return unit, err
	}

	err2 := u.C.Bind(&unit)
	if err2 != nil {
		u.C.Logger().Error(err2.Error)
		return unit, err2
	}

	err3 := u.DB.Model(&unit).Where("id = ?", ids).Updates(&unit)
	if err3.Error != nil {
		u.C.Logger().Error(err3.Error)
		return unit, err3.Error
	}

	var updatedUnit models.WorkUnit
	err4 := u.DB.Find(&updatedUnit, ids)
	if err4.Error != nil {
		u.C.Logger().Error(err4.Error)
		return unit, err4.Error
	}
	
	return updatedUnit, nil
}

func (u *UnitServices) Delete() (models.WorkUnit, error) {
	var unit models.WorkUnit
	id := u.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		u.C.Logger().Error(err.Error)
		return unit, err
	}
	
	err2 := u.DB.Delete(&unit, ids)
	if err2.Error != nil {
		u.C.Logger().Error(err2.Error)
		return unit, err2.Error
	}
	
	return unit, nil
}