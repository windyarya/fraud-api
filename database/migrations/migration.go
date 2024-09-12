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
	)

	if err != nil {
		db.Rollback()
	}
}