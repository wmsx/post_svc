package models

type PostItem struct {
	Model
	PostId   int64
	ObjectId int64
	Index    int32
	Status   byte
}

func (PostItem) TableName() string  {
	return "t_post_item"
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
