package models

type CollectionItem struct {
	Model
	CollectionId int64
	PostId       int64
	feature      byte
	Status       byte
}
