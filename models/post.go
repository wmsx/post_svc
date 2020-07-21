package models

import "gorm.io/gorm"

type Post struct {
	Model
	Title       string
	Description string
	Type        int32
	CategoryId  int64
	MengerId    int64
	Status      byte
}

func (Post) TableName() string  {
	return "t_post"
}

func GetPostList(categoryId, lastId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("category_id = ? and id > ?", categoryId, lastId).Limit(10).Find(&posts).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return posts, nil
}

func AddPost(post *Post) error {
	err := db.Create(post).Error
	if err != nil {
		return err
	}
	return nil
}
