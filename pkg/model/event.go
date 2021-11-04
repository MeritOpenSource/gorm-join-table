package model

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Name string
}
