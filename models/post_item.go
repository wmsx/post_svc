package models

import "github.com/jinzhu/gorm"

type PostItem struct {
	gorm.Model
	PostId   int64
	ObjectId int64
	Index    int32
	Status   byte
}

func GetItemByPostIds(postIds []int64) ([]*PostItem, error) {
	var items []*PostItem
	err := db.Where("post_id in (?)", postIds).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func BatchAddPostItem(postItems []*PostItem) error {
	err := db.Create(&postItems).Error
	if err != nil {
		return err
	}
	return nil
}
