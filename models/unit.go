package models

type WorkUnit struct {
	GormModel
	Name string `json:"name"`
	Address string `json:"address"`
	Type string `json:"type"`
	Description string `json:"description"`
}