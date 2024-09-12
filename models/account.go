package models

type Account struct {
    GormModel
    Number              string           `json:"number"`
    Balance             string    	     `json:"balance"`
    AverageTrx          string           `json:"average_trx"`
    Currency            string           `json:"currency"`
    AccountStatusID     uint             `json:"status_id"`
    AccountIdentityID   uint             `json:"identity_id"`
    WorkUnitID          uint             `json:"work_unit_id"`
    WorkUnit            WorkUnit         `gorm:"foreignKey:WorkUnitID"`
    AccountIdentity     AccountIdentity  `gorm:"foreignKey:AccountIdentityID"`
    AccountStatus       AccountStatus    `gorm:"foreignKey:AccountStatusID"`
}

type AccountIdentity struct {
    GormModel
    Name      string `json:"name"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Password  string `json:"password"`
    NIK       string `json:"nik"`
}

type AccountStatus struct {
    GormModel
    Status string `json:"status"`
}

type AccountRequest struct {
	Name      string `json:"name"`
    Email     string `json:"email"`
    Phone     string `json:"phone"`
    Password  string `json:"password"`
    NIK       string `json:"nik"`
	Number              string           `json:"number"`
    Balance             string    	     `json:"balance"`
    Currency            string           `json:"currency"`
    AccountStatusID     uint             `json:"status_id"`
    AccountIdentityID   uint             `json:"identity_id"`
    WorkUnitID          uint             `json:"work_unit_id"`
    WorkUnit            WorkUnit         `gorm:"foreignKey:WorkUnitID"`
    AccountIdentity     AccountIdentity  `gorm:"foreignKey:AccountIdentityID"`
    AccountStatus       AccountStatus    `gorm:"foreignKey:AccountStatusID"`
}