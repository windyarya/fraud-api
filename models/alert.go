package models

import "time"

type Alert struct {
	GormModel
	Name string `json:"name"`
	Description string `json:"description"`
	Comment string `json:"comment"`
	AlertStatusID uint `json:"status_id"`
	ActivityID uint `json:"activity_id"`
	AlertStatus AlertStatus `gorm:"foreignKey:AlertStatusID"`
	Activity Activity `gorm:"foreignKey:ActivityID"`
	Users []User `gorm:"many2many:alert_users;"`
}

type AlertStatus struct {
	GormModel
	Status string `json:"status"`
}

type AlertLog struct {
	GormModel
	AlertID       int      `json:"alert_id"`
	Alert         Alert     `gorm:"foreignKey:AlertID"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	AlertStatusID uint      `json:"status_id"`
	ActivityID    uint      `json:"activity_id"`
	ChangedBy     uint      `json:"changed_by"`
	ChangeReason  string    `json:"change_reason"`
}

type AlertUser struct {
	AlertID uint `gorm:"primaryKey"`
	UserID  uint `gorm:"primaryKey"`
	CreatedAt time.Time
}
