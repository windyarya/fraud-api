package models

type UserGroup struct {
	GormModel
	Name string `json:"name"`
	Description string `json:"description"`
}