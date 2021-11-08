package join

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
	err = create(db)
	if err != nil {
		return err
	}
	fmt.Println("join db insert successful")
	if err != nil {
		return err
	}
	err = query(db)
	fmt.Println("join db query successful")
	return nil
}

func test(db *gorm.DB) error {
	return validate.All(db, []interface{}{Kiosk{}, Event{}, KioskEvent{}} )
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

func create(db *gorm.DB) error {
	err  := createEvent(db, "Exciting Activity!")
	if err != nil {
		return err
	}
	err = createKiosk(db, "First Kiosk")
	if err != nil {
		return err
	}
	err = createKiosk(db, "Another Kiosk")
	if err != nil {
		return err
	}
	err = linkKioskAndEvent(1, 1, db)
	if err != nil {
		return err
	}
	return nil
}

func getEvent(eventID uint, db *gorm.DB) (Event, error) {
	event := Event{ Model: gorm.Model{ID: eventID}}
	result := db.Preload("Kiosks").First(&event)
	return event, result.Error
}

func getKiosk( kioskID uint, db *gorm.DB) (Kiosk, error) {
	kiosk := Kiosk{ Model: gorm.Model{ID: kioskID}}
	result := db.Preload("Events").First(&kiosk)
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
