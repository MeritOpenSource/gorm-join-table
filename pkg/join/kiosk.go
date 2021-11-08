package join

import "gorm.io/gorm"

type Kiosk struct {
	gorm.Model
	Events          []Event `gorm:"many2many:kiosk_events;"`
	Name string
}
