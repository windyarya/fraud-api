package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type GroupServices struct {
	DB *gorm.DB
	C echo.Context
}

func (g *GroupServices) GetAll() ([]models.UserGroup, error) {
	var groups []models.UserGroup

	err := g.DB.Find(&groups)
	println(err)
	if err.Error != nil {
		println(err)
		return groups, err.Error
	}
	
	return groups, nil
}

func (g *GroupServices) GetByID() (models.UserGroup, error) {
	var group models.UserGroup
	id := g.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		g.C.Logger().Error(err.Error)
		return group, err
	}
	
	err2 := g.DB.Find(&group, ids)
	if err2.Error != nil {
		g.C.Logger().Error(err2.Error)
		return group, err2.Error
	}
	
	return group, nil
}

func (g *GroupServices) Create() (models.UserGroup, error) {
	var group models.UserGroup
	err := g.C.Bind(&group)
	if err != nil {
		g.C.Logger().Error(err.Error)
		return group, err
	}
	
	err2 := g.DB.Create(&group)
	if err2.Error != nil {
		g.C.Logger().Error(err2.Error)
		return group, err2.Error
	}
	
	return group, nil
}

func (g *GroupServices) Update() (models.UserGroup, error) {
	var group models.UserGroup

	id := g.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		g.C.Logger().Error(err.Error)
		return group, err
	}

	err2 := g.C.Bind(&group)
	if err2 != nil {
		g.C.Logger().Error(err2.Error)
		return group, err2
	}

	err3 := g.DB.Model(&group).Where("id = ?", ids).Updates(&group)
	if err3.Error != nil {
		g.C.Logger().Error(err3.Error)
		return group, err3.Error
	}

	var updatedGroup models.UserGroup
	err4 := g.DB.Find(&updatedGroup, ids)
	if err4.Error != nil {
		g.C.Logger().Error(err4.Error)
		return group, err4.Error
	}
	
	return updatedGroup, nil
}

func (g *GroupServices) Delete() (models.UserGroup, error) {
	var group models.UserGroup
	id := g.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		g.C.Logger().Error(err.Error)
		return group, err
	}
	
	err2 := g.DB.Delete(&group, ids)
	if err2.Error != nil {
		g.C.Logger().Error(err2.Error)
		return group, err2.Error
	}
	
	return group, nil
}