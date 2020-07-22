package models

type Tag struct {
	Model
	PostId int64  `gorm:"comment:关联的postId"`
	Name   string `gorm:"comment:名称"`
}

