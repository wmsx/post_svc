package models

import "github.com/jinzhu/gorm"

type CollectionItem struct {
	gorm.Model
	CollectionId int64
	PostId       int64
	feature      byte
	Status       byte
}
