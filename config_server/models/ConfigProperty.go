package models

import "gorm.io/gorm"

type ConfigProperty struct {
	gorm.Model
	ConfigID  uint   `json:"config_id"`
	SectionID uint   `json:"section_id"`
	SKey      string `json:"Skey"`
	Value     string `json:"Value"`
}
