package models

type PostItem struct {
	Model
	PostId   int64 `gorm:"index:idx_post_id;comment:关联post的ID"`
	ObjectId int64 `gorm:"comment:对象id"`
	Index    int32 `gorm:"comment:在整个post的序号"`
	Type     int32 `gorm:"comment:类型 1-图片 2-视频"`
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
