package models

import "gorm.io/gorm"

type ConfigProperty struct {
	gorm.Model
	ConfigID  uint   `json:"config_id"`
	SectionID uint   `json:"section_id"`
	Key       string `json:"key"`
	Value     string `json:"Value"`
}
