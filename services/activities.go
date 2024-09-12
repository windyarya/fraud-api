package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"bitbucket.org/windyarya/backend-final/services/apis"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ActivityServices struct {
	DB *gorm.DB
	C  echo.Context
}

func (a *ActivityServices) GetAll() ([]models.Activity, error) {
	var activities []models.Activity
	err := a.DB.Preload("Account").Find(&activities)
	if err.Error != nil {
		return activities, err.Error
	}
	return activities, nil
}

func (a *ActivityServices) GetByID() (models.Activity, error) {
	var activity models.Activity
	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return activity, err
	}

	err2 := a.DB.Preload("Account").Find(&activity, ids)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return activity, err2.Error
	}

	return activity, nil
}

func (a *ActivityServices) Create() (models.Activity, error) {
	var activity models.Activity
	err := a.C.Bind(&activity)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return activity, err
	}

	var account models.Account
	id := activity.AccountID

	err3 := a.DB.Find(&account, id)
	if err3.Error != nil {
		a.C.Logger().Error(err3.Error)
		return activity, err3.Error
	}

	avgTrx, err4 := strconv.ParseFloat(account.AverageTrx, 64)
	if err4 != nil {
		a.C.Logger().Error(err4.Error)
		return activity, err4
	}

	trxAmount, err5 := strconv.ParseFloat(activity.Amount, 64)
	if err5 != nil {
		a.C.Logger().Error(err5.Error)
		return activity, err5
	}

	if trxAmount > (3 * avgTrx) {
		activity.Flag = true
		activity.Severity = "High"
	} else {
		activity.Flag = false
		activity.Severity = "Low"
	}

	err2 := a.DB.Create(&activity)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return activity, err2.Error
	}

	err7 := a.DB.Preload("Account").Find(&activity, activity.ID)
	if err7.Error != nil {
		a.C.Logger().Error(err7.Error)
		return activity, err7.Error
	}

	if activity.Flag {
		alert := models.Alert{
			ActivityID:    activity.ID,
			Name:          "#" + strconv.Itoa(int(activity.ID)) + " - Suspicious High Transaction Amount Detected",
			Description:   "High and not usual transaction amount detected",
			AlertStatusID: 1,
		}
	
		err6 := a.DB.Create(&alert)
		if err6.Error != nil {
			a.C.Logger().Error(err6.Error)
			return activity, err6.Error
		}
	
		webhookURL := "https://discord.com/api/webhooks/1283797408630706309/PaqHqIra0bOJO_2MlXWFpEl5kkcmqbt-yJJBGUE9kbwdia5sA43G_0ixU0cJ8GBibtE5"
		if err := apis.SendNotification(webhookURL, alert); err != nil {
			a.C.Logger().Error("Failed to send notification to Discord: " + err.Error())
		}
	}

	return activity, nil
}
