package kiosk

import (
	"fmt"
	"join_table/pkg/validate"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	err := test(db)
	if err != nil {
		return err
	}
	fmt.Println("join DB Validation successful")

	err = query(db)
	if err != nil {
		return err
	}
	fmt.Println("successfully queried for an kiosk's events and all the checkins for those events")
	return nil
}

func test(db *gorm.DB) error {
	return validate.All(db, []interface{}{Kiosk{}, Event{}, KioskEvent{}})
}

func query(db *gorm.DB) error {
	fmt.Println("Kiosks")
	kiosk, err := getKiosk(1, db)
	if err != nil {
		return err
	}
	fmt.Println(kiosk)

	fmt.Println("Event")
	event, err := getEvent(1, db)
	if err != nil {
		return err
	}
	fmt.Println(event)
	return nil
}

func getEvent(eventID uint, db *gorm.DB) (Event, error) {
	event := Event{Model: gorm.Model{ID: eventID}}
	result := db.Preload("Checkins").First(&event)
	return event, result.Error
}

func getKiosk(kioskID uint, db *gorm.DB) (Kiosk, error) {
	kiosk := Kiosk{Model: gorm.Model{ID: kioskID}}
	result := db.Preload("Events.Checkins").First(&kiosk)
	return kiosk, result.Error
}

func linkKioskAndEvent(kioskID, eventID uint, db *gorm.DB) error {
	kioskEvent := KioskEvent{
		KioskID: kioskID,
		EventID: eventID,
	}
	result := db.Create(&kioskEvent)
	return result.Error
}

func createKiosk(db *gorm.DB, name string) error {
	kiosk := Kiosk{
		Name: name,
	}
	result := db.Create(&kiosk)
	return result.Error
}

func createEvent(db *gorm.DB, name string) error {
	event := Event{
		Name: name,
	}
	result := db.Create(&event)
	return result.Error
}
