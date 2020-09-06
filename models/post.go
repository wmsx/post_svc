package models

import "gorm.io/gorm"

type Post struct {
	Model
	Title       string `gorm:"comment:标题"`
	Description string `gorm:"comment:描述"`
	Type        int32  `gorm:"comment:类型 0-序列集 1-图片 2-视频"`
	CategoryId  int64  `gorm:"comment:所属分类"`
	MengerId    int64  `gorm:"comment:创建人"`
}

func GetPostList(categoryId, lastId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("category_id = ? and id > ?", categoryId, lastId).Limit(10).Find(&posts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return posts, nil
}

func GetPostByIds(postIds []int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("id in (?)", postIds).Find(&posts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return posts, nil
}

func GetMengerPostList(mengerId int64, pageNum, pageSize int32) ([]*Post, error) {

	var id int64
	db.Table("t_post").
		Select("id").
		Where("menger_id = ?", mengerId).
		Offset(int(pageNum * pageSize)).
		Limit(1).
		Scan(&id)

	var posts []*Post

	err := db.Where("id > ?", id).Limit(int(pageSize)).Find(&posts).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
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
