package models

import (
	"gorm.io/gorm"
)

type GoldPrice struct {
	gorm.Model
	GDate   string
	GOpen   float32
	GHigh   float32
	GLow    float32
	GClose  float32
	GVolume float32
}
