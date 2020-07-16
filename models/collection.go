package models

import "github.com/jinzhu/gorm"

type Collection struct {
	gorm.Model
	Name       string
	Cover      string
	CategoryId int64
	MengerId   int64
	Status     byte
}

