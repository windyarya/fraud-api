package models

type Activity struct {
	GormModel
	Amount string `json:"amount"`
	Currency string `json:"currency"`
	Flag bool `json:"flag"`
	Severity string `json:"severity"`
	Location string `json:"location"`
	IPAddress string `json:"ip_address"`
	AccountID uint `json:"account_id"`
	Account Account `gorm:"foreignKey:AccountID"`
}