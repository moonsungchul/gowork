package models

import "gorm.io/gorm"

type ConfigSection struct {
	gorm.Model
	ConfigID uint   `json:"config_id"`
	Section  string `json:"section"`
}
