package handler

import (
	"context"
	"github.com/wmsx/post_svc/models"
	proto "github.com/wmsx/post_svc/proto/category"
)

type CategoryHandler struct{}

// 获取所有分类
func (c CategoryHandler) GetAllCategory(ctx context.Context,
	req *proto.GetAllCategoryRequest, res *proto.GetAllCategoryResponse) error {
	var (
		err error
		categories []*models.Category
	)

	if categories, err = models.GetAllCategory(); err != nil {
		return err
	}

	categoryInfos := make([]*proto.CategoryInfo, 0)
	for _, category := range categories{
		categoryInfo := &proto.CategoryInfo{
			Id:       int64(category.ID),
			Name:     category.Name,
			MengerId: category.MengerId,
		}
		categoryInfos = append(categoryInfos, categoryInfo)
	}

	res.CategoryInfos = categoryInfos
	return nil
}

// 创建分类
func (c CategoryHandler) CreateCategory(ctx context.Context,
	req *proto.CreateCategoryRequest, res *proto.CreateCategoryResponse) error {
	
	category := &models.Category{
		Name:   req.Name,
		Status: 1,
	}
	if err := models.CreateCategory(category); err != nil  {
		return err
	}
	res = &proto.CreateCategoryResponse{}
	return nil
}
