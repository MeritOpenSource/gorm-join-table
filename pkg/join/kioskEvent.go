package join

type KioskEvent struct {
	KioskID  uint       `gorm:"primaryKey"`
	EventID  uint       `gorm:"primarykey"`
}
