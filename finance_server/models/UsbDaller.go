package models

import "gorm.io/gorm"

type UsbEur struct {
	gorm.Model
	GDate   string  `json:"gdate"`
	GClose  float32 `json:"close"`
	GOpen   float32 `json:"open"`
	GHigh   float32 `json:"high"`
	GLow    float32 `json:"low"`
	GChange float32 `json:"change"`
}

type UsbKrw struct {
	gorm.Model
	GDate   string  `json:"gdate"`
	GClose  float32 `json:"close"`
	GOpen   float32 `json:"open"`
	GHigh   float32 `json:"high"`
	GLow    float32 `json:"low"`
	GChange float32 `json:"change"`
}

type UsbJpy struct {
	gorm.Model
	GDate   string  `json:"gdate"`
	GClose  float32 `json:"close"`
	GOpen   float32 `json:"open"`
	GHigh   float32 `json:"high"`
	GLow    float32 `json:"low"`
	GChange float32 `json:"change"`
}

type UsbRub struct {
	gorm.Model
	GDate   string  `json:"gdate"`
	GClose  float32 `json:"close"`
	GOpen   float32 `json:"open"`
	GHigh   float32 `json:"high"`
	GLow    float32 `json:"low"`
	GChange float32 `json:"change"`
}
