package main

import (
	"fmt"
	"join_table/pkg/checkin"
	"join_table/pkg/join"
	"join_table/pkg/kiosk"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting")
	format := "host=%s user=%s password=%s dbname=%s port=%d search_path=%s sslmode=disable TimeZone=America/Los_Angeles"
	dsn := fmt.Sprintf(format, "localhost", "postgres", "password", "postgres", 5432, "checkin")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	err = join.Run(db)
	if err != nil {
		panic(err)
	}

	err = checkin.Run(db)
	if err != nil {
		panic(err)
	}

	err = kiosk.Run(db)
	if err != nil {
		panic(err)
	}
}
