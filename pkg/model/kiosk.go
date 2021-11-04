package model

import "gorm.io/gorm"

type Kiosk struct {
	gorm.Model
	Name string
}
