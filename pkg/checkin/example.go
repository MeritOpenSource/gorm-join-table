package checkin

import (
	"errors"
	"fmt"
	"time"

	"join_table/pkg/validate"

	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	err := test(db)
	if err != nil {
		return err
	}
	fmt.Println("checkin db validation successful!")
	err = create(db)
	if err != nil {
		return err
	}
	fmt.Println("checkin db insert successful")
	err = query(db)
	if err != nil {
		return err
	}
	fmt.Println("checkin db query successful")
	return nil
}

func test(db *gorm.DB) error {
	return validate.All(db, []interface{}{KioskEvent{}, Checkin{}})
}

func query(db *gorm.DB) error {
	kioskEvent, err := getKioskEvent(1, 1, db)
	if err != nil {
		return err
	}
	fmt.Println("Checkins")
	fmt.Println(kioskEvent.Checkins)
	return nil
}

func create(db *gorm.DB) error {
	err := createCheckin(db, 1, 1, "first checkin")
	if err != nil {
		return err
	}
	err = createCheckin(db, 1, 1, "second checkin")
	if err != nil {
		return err
	}
	err = addBadCheckin(db)
	if err != nil {
		return err
	}
	return nil
}

func getKioskEvent(kioskID, eventID uint, db *gorm.DB) (KioskEvent, error) {
	kioskEvent := KioskEvent{}
	result := db.Preload("Checkins").Where(&KioskEvent{KioskID: kioskID, EventID: eventID}).First(&kioskEvent)
	return kioskEvent, result.Error
}

func createCheckin(db *gorm.DB, eventID, kioskID uint, name string) error {
	checkin := Checkin{
		EventID:         eventID,
		KioskID:         kioskID,
		CheckinDatetime: time.Now(),
		Name:            name,
	}
	result := db.Create(&checkin)
	return result.Error
}

// addBadCheckin gorm automatically logs calls that fail.
// This func is designed to fail, so calling it _should_ be noisy
func addBadCheckin(db *gorm.DB) error {
	invalidCheckin := Checkin{
		EventID:         1,
		KioskID:         2,
		CheckinDatetime: time.Time{},
		Name:            "second checkin",
	}
	fmt.Print("error expected for an insert on table \"checkins\"")
	result := db.Create(&invalidCheckin)
	if result.Error == nil {
		return errors.New("expected an error when create a checkin against an invalid event/kiosk pair")
	}
	return nil
}
