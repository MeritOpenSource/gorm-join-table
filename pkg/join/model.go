package join

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Kiosks []Kiosk `gorm:"many2many:kiosk_events;"`
	Name   string
}

type Kiosk struct {
	gorm.Model
	Events []Event `gorm:"many2many:kiosk_events;"`
	Name   string
}

type KioskEvent struct {
	KioskID uint `gorm:"primaryKey"`
	EventID uint `gorm:"primarykey"`
}
