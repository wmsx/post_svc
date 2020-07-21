package models

type Tag struct {
	Model
	PostId int64
	Name   string
	Status byte
}

func (Tag) TableName() string {
	return "t_tag"
}

