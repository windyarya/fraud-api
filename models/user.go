package models

type User struct {	
	GormModel
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
    UserGroupID   uint       `json:"user_group_id"`
    UserGroup     UserGroup  `gorm:"foreignKey:UserGroupID"`
    WorkUnitID    uint       `json:"work_unit_id"`
    WorkUnit      WorkUnit   `gorm:"foreignKey:WorkUnitID"`
	Alerts []Alert   `gorm:"many2many:alert_users;"`
}

type UserResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	UserGroupID   uint       `json:"user_group_id"`
	WorkUnitID    uint       `json:"work_unit_id"`
}
