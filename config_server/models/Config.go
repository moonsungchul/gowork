package models

import "gorm.io/gorm"

type Config struct {
	gorm.Model
	CName string `json:"cname"`
}
