package models

import (
	"gorm.io/gorm"
)

type GoldPrice struct {
	gorm.Model
	GDate   string  `json:"gdate"`
	GOpen   float32 `json:"gopen"`
	GHigh   float32 `json:"ghigh"`
	GLow    float32 `json:"glow"`
	GClose  float32 `json:"gclose"`
	GVolume float32 `json:"gvolume"`
}
