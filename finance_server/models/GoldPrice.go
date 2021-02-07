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

	CloseHigh         float32 `json:"CloseHigh"`
	OpenClose         float32 `json:"OpenClose"`
	LowClose          float32 `json:"LowClose"`
	Close5LastClose   float32 `json:"Close5LastClose"`
	Close10LastClose  float32 `json:"Close10LastClose"`
	Close20LastClose  float32 `json:"Close20LastClose"`
	Close60LastClose  float32 `json:"Close60LastClose"`
	Close120LastClose float32 `json:"Close120LastClose"`

	Volume5LastVolume   float32 `json:"Volume5LastVolume"`
	Volume10LastVolume  float32 `json:"Volume10LastVolume"`
	Volume20LastVolume  float32 `json:"Volume20LastVolume"`
	Volume60LastVolume  float32 `json:"Volume60LastVolume"`
	Volume120LastVolume float32 `json:"Volume120LastVolume"`
}
