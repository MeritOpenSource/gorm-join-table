package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Kiosks []Kiosk `gorm:"many2many:kiosk_events;"`
	Name string
}
