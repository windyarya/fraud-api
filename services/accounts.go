package services

import (
	"bitbucket.org/windyarya/backend-final/models"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AccountServices struct {
	DB *gorm.DB
	C echo.Context
}

func (a *AccountServices) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	err := a.DB.Preload("AccountIdentity").Preload("AccountStatus").Preload("WorkUnit").Find(&accounts)
	if err.Error != nil {
		return accounts, err.Error
	}
	return accounts, nil
}

func (a *AccountServices) GetByID() (models.Account, error) {
	var account models.Account
	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return account, err
	}
	
	err2 := a.DB.Preload("AccountIdentity").Preload("AccountStatus").Preload("WorkUnit").Find(&account, ids)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return account, err2.Error
	}
	
	return account, nil
}

func (a *AccountServices) CreateIdentity(identity models.AccountIdentity) (error) {
	err := a.DB.Create(&identity)
	if err.Error != nil {
		a.C.Logger().Error(err.Error)
		return err.Error
	}
	
	return nil
}

func (a *AccountServices) Create() (models.AccountRequest, error) {
	var request models.AccountRequest

	err := a.C.Bind(&request)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return request, err
	}

	identity := models.AccountIdentity{
        Name:     request.Name,
        Email:    request.Email,
        Phone:    request.Phone,
        Password: request.Password,
        NIK:      request.NIK,
    }

	var identityCheck models.AccountIdentity
	err4 := a.DB.Find(&identityCheck, "nik = ?", identity.NIK)
	if err4.Error != nil && identityCheck.NIK != identity.NIK {
		err2 := a.CreateIdentity(identity)
		if err2 != nil {
			a.C.Logger().Error(err2.Error)
			return request, err2
		}	
	}

	err5 := a.DB.First(&identityCheck, "nik = ?", identity.NIK)
	if err5.Error != nil {
		a.C.Logger().Error(err5.Error)
		return request, err5.Error
	}
	identity.ID = identityCheck.ID

	account := models.Account{
        Number:            request.Number,
        Balance:           request.Balance,
        Currency:          request.Currency,
		AverageTrx:        request.AverageTrx,
        AccountStatusID:   request.AccountStatusID,
        WorkUnitID:        request.WorkUnitID,
        AccountIdentityID: identityCheck.ID,
    }

	err3 := a.DB.Create(&account)
	if err3.Error != nil {
		a.C.Logger().Error(err3.Error)
		return request, err3.Error
	}

	err5 := a.DB.Preload("AccountIdentity").Preload("AccountStatus").Preload("WorkUnit").Find(&account, account.ID)
	if err5.Error != nil {
		a.C.Logger().Error(err5.Error)
		return request, err5.Error
	}
	
	return request, nil
}

func (a *AccountServices) Update() (models.Account, error) {
	var request models.AccountRequest
	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
	}

	err2 := a.C.Bind(&request)
	account := models.Account{
		Number:            request.Number,
		Balance:           request.Balance,
		Currency:          request.Currency,
		AccountStatusID:   request.AccountStatusID,
		WorkUnitID:        request.WorkUnitID,
		AccountIdentityID: request.AccountIdentityID,
	}
	if err2 != nil {
		a.C.Logger().Error(err2.Error)
		return account, err2
	}

	if account != (models.Account{}) {
		err := a.DB.Model(&account).Where("id = ?", ids).Updates(&account)
		if err.Error != nil {
			a.C.Logger().Error(err.Error)
			return account, err.Error
		}
	}

	identity := models.AccountIdentity{
		Name:     request.Name,
		Email:    request.Email,
		Phone:    request.Phone,
		Password: request.Password,
		NIK:      request.NIK,
	}

	if identity != (models.AccountIdentity{}) {
		var account2 models.Account
		err1 := a.DB.Find(&account2, ids)
		if err1.Error != nil {
			a.C.Logger().Error(err1.Error)
			return account, err1.Error
		}
		err := a.DB.Model(&identity).Where("id = ?", account2.AccountIdentityID).Updates(&identity)
		if err.Error != nil {
			a.C.Logger().Error(err.Error)
			return account, err.Error
		}
	}

	var updatedAccount models.Account
	err4 := a.DB.Find(&updatedAccount, ids)
	if err4.Error != nil {
		a.C.Logger().Error(err4.Error)
		return account, err4.Error
	}

	return updatedAccount, nil
}

func (a *AccountServices) Delete() (models.Account, error) {
	var account models.Account

	id := a.C.Param("id")
	ids, err := strconv.Atoi(id)
	if err != nil {
		a.C.Logger().Error(err.Error)
		return account, err
	}

	err2 := a.DB.Delete(&account, ids)
	if err2.Error != nil {
		a.C.Logger().Error(err2.Error)
		return account, err2.Error
	}
	
	return account, nil
}