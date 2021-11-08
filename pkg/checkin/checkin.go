package checkin

import (
	"time"

	"gorm.io/gorm"
)

type Checkin struct {
	gorm.Model
	EventID         uint
	KioskID         uint
	CheckinDatetime time.Time
	Name            string
}
