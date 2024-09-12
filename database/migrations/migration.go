package migration

import (
	"bitbucket.org/windyarya/backend-final/models"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	err := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.User{}, 
		&models.UserGroup{},
		&models.WorkUnit{},
		&models.Account{},
		&models.AccountIdentity{},
		&models.AccountStatus{},
	)

	if err != nil {
		db.Rollback()
	}
}