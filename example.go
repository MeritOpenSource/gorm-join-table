package main

import (
	"fmt"

	"join_table/pkg/model"
	"join_table/pkg/validate"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting")
	format := "host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=disable TimeZone=America/Los_Angeles"
	dsn := fmt.Sprintf(format, "localhost", "postgres", "password", "postgres", 5432, "checkin")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("DB connection successful")
	err = validate.All(db) // do all this to test that gorm is happy with our liquibase file
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Validation successful")

	createEvent(db)
	createKiosk(db)
	linkKioskAndEvent(1, 1, db)

	kioskEvent := getKioskEvent(1, 1, db)
	fmt.Println("KioskEvent")
	fmt.Println(kioskEvent)
	fmt.Println("Kiosks")
	fmt.Println(getKiosk(1, db))
	fmt.Println("Event")
	fmt.Println(getEvent(1, db))
}

func getEvent(eventID uint, db *gorm.DB) model.Event {
	event := model.Event{ Model: gorm.Model{ID: eventID}}
	result := db.Preload("Kiosks").First(&event)
	if result.Error != nil {
		panic(result.Error)
	}
	return event
}

func getKiosk( kioskID uint, db *gorm.DB) model.Kiosk {
	kiosk := model.Kiosk{ Model: gorm.Model{ID: kioskID}}
	result := db.Preload("Events").First(&kiosk)
	if result.Error != nil {
		panic(result.Error)
	}
	return kiosk
}

func getKioskEvent(kioskID, eventID uint, db *gorm.DB) model.KioskEvent {
	kioskEvent := model.KioskEvent{}
	result := db.Where(&model.KioskEvent{KioskID: kioskID, EventID: eventID}).First(&kioskEvent)
	if result.Error != nil {
		panic(result.Error)
	}
	return kioskEvent
}

func linkKioskAndEvent(kioskID, eventID uint, db *gorm.DB) {
	kioskEvent := model.KioskEvent{
		KioskID: kioskID,
		EventID: eventID,
	}
	result := db.Create(&kioskEvent)
	if result.Error != nil {
		panic(result.Error)
	}
}

func createKiosk(db *gorm.DB) {
	firstKiosk := model.Kiosk{
		Name: "Checkin Kiosk",
	}
	result := db.Create(&firstKiosk)
	if result.Error != nil {
		panic(result.Error)
	}
	secondKiosk := model.Kiosk{
		Name: "Another kiosk",
	}
	result = db.Create(&secondKiosk)
	if result.Error != nil {
		panic(result.Error)
	}
}

func createEvent(db *gorm.DB) {
	event := model.Event{
		Name: "Exciting Activity!",
	}
	result := db.Create(&event)
	if result.Error != nil {
		panic(result.Error)
	}
}
