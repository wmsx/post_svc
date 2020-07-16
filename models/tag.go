package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	PostId int64
	Name string
	Status byte
}
