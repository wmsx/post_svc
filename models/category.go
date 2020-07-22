package models

import "gorm.io/gorm"

type Category struct {
	Model
	Name     string `gorm:"not null;type:varchar(16);comment:名字"`
	ShowName string `gorm:"not null;type:varchar(10);comment:显示名"`
	MengerId int64  `gorm:"not null;comment:创建人"`
}

func GetAllCategory() ([]*Category, error) {
	var categories []*Category
	err := db.Find(&categories).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return categories, nil
}

func CreateCategory(category *Category) error {
	err := db.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}
