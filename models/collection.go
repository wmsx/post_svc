package models

type Collection struct {
	Model
	Name       string
	Cover      string
	CategoryId int64
	MengerId   int64
	Status     byte
}

