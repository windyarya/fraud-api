package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AlertServices struct {
	DB *gorm.DB
	C echo.Context
}

func (a *AlertServices) GetAll() ([]models.Alert, error) {
	var alerts []models.Alert
	err := a.DB.Find(&alerts)
	if err.Error != nil {
		return alerts, err.Error
	}
	return alerts, nil
}

func (a *AlertServices) GetByID() (models.Alert, error) {
	var alert models.Alert
	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return alert, err
	}
	
	err2 := a.DB.Find(&alert, ids)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return alert, err2.Error
	}
	
	return alert, nil
}

func (a *AlertServices) Create() (models.Alert, error) {
	var alert models.Alert
	err := a.C.Bind(&alert)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return alert, err
	}
	
	err2 := a.DB.Create(&alert)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return alert, err2.Error
	}
	
	return alert, nil
}

func (a *AlertServices) CreateLog(log models.AlertLog) error {
	err := a.DB.Create(&log)
	if err.Error != nil {
		a.C.Logger().Error(err.Error)
		return err.Error
	}

	return nil
}

func (a *AlertServices) Update() (models.Alert, error) {
	var alert models.Alert
	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return alert, err
	}

	err2 := a.C.Bind(&alert)
	if err2 != nil {
		a.C.Logger().Error(err2.Error)
		return alert, err2
	}

	err3 := a.DB.Model(&alert).Where("id = ?", ids).Updates(&alert)
	if err3.Error != nil {
		a.C.Logger().Error(err3.Error)
		return alert, err3.Error
	}

	var alert_user models.AlertUser
	err5 := a.DB.Find(&alert_user, "alert_id = ?", ids)
	if err5.Error != nil {
		a.C.Logger().Error(err5.Error)
		return alert, err5.Error
	}

	log := models.AlertLog{
		AlertID: ids,
		Name: alert.Name,   
		Description: alert.Description,
		AlertStatusID: alert.AlertStatusID,
		ActivityID: alert.ActivityID,
		ChangedBy: alert_user.UserID,
		ChangeReason: "",
	}

	err6 := a.CreateLog(log)
	if err6 != nil {
		a.C.Logger().Error(err6.Error)
		return alert, err6
	}

	var updatedAlert models.Alert
	err4 := a.DB.Find(&updatedAlert, ids)
	if err4.Error != nil {
		a.C.Logger().Error(err4.Error)
		return alert, err4.Error
	}

	return updatedAlert, nil
}

func (a *AlertServices) Delete() (models.Alert, error) {
	var alert models.Alert

	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return alert, err
	}
	
	err2 := a.DB.Delete(&alert, ids)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return alert, err2.Error
	}
	
	return alert, nil
}