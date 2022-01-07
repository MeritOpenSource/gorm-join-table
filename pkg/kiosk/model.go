package kiosk

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	// Kiosks      []Kiosk `gorm:"many2many:kiosk_events"`
	// KioskEvents []KioskEvent
	Checkins []Checkin `gorm:"foreignKey:event_id"`
	Name     string
}

type Kiosk struct {
	gorm.Model
	Events []Event `gorm:"many2many:kiosk_events"`
	Name   string
}

type Checkin struct {
	gorm.Model
	EventID          uint
	KioskID          uint
	CheckinDatetime  time.Time
	PlatformMemberID string
}

type KioskEvent struct {
	KioskID  uint      `gorm:"primaryKey"`
	EventID  uint      `gorm:"primarykey"`
	Checkins []Checkin `gorm:"foreignKey:KioskID,EventID;References:KioskID,EventID"`
}
