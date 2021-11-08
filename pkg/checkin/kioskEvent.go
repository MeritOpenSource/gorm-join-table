package checkin

type KioskEvent struct {
	Checkins []Checkin `gorm:"foreignKey:KioskID,EventID;References:KioskID,EventID"`
	KioskID  uint       `gorm:"primaryKey"`
	EventID  uint       `gorm:"primarykey"`
}
