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

type KioskEvent struct {
	Checkins []Checkin `gorm:"foreignKey:KioskID,EventID;References:KioskID,EventID"`
	KioskID  uint      `gorm:"primaryKey"`
	EventID  uint      `gorm:"primarykey"`
}
