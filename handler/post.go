package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/post_svc/models"
	proto "github.com/wmsx/post_svc/proto/post"
)

type PostHandler struct{}

func (p *PostHandler) SavePost(ctx context.Context, req *proto.CreatePostRequest, res *proto.CreatePostResponse) error {
	post := &models.Post{
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		CategoryId:  req.CategoryId,
		MengerId:    req.MengerId,
		Status:      1,
	}
	_ = models.AddPost(post)
	postItems := make([]*models.PostItem, 0)
	for _, protoPostItem := range req.Items {
		postItem := &models.PostItem{
			PostId:   post.ID,
			ObjectId: protoPostItem.ObjectId,
			Index:    protoPostItem.Index,
			Status:   1,
		}
		postItems = append(postItems, postItem)
	}
	_ = models.BatchAddPostItem(postItems)
	return nil
}

func (p *PostHandler) GetPostList(ctx context.Context, req *proto.GetPostListRequest, res *proto.GetPostListResponse) error {
	var (
		posts     []*models.Post
		postItems []*models.PostItem
		err       error
	)
	if posts, err = models.GetPostList(req.CategoryId, req.LastId); err != nil {
		log.Error("从数据库获取Post列表失败, err: ", err)
		return err
	}

	postIds := make([]int64, 0)
	for _, post := range posts {
		postIds = append(postIds, int64(post.ID))
	}

	if postItems, err = models.GetItemByPostIds(postIds); err != nil {
		log.Error("从数据库获取PostItem列表失败, err: ", err)
		return err
	}

	post2ItemMap := make(map[int64][]*models.PostItem)
	for _, postItem := range postItems {
		var items []*models.PostItem
		var ok bool
		if items, ok = post2ItemMap[postItem.PostId]; !ok {
			items = make([]*models.PostItem, 0)
			post2ItemMap[postItem.PostId] = items
		}
		items = append(items, postItem)
	}

	postInfos := make([]*proto.PostInfo, 0)
	for _, post := range posts {
		protoPostItems := make([]*proto.PostItem, 0)
		if modelPostItems, ok := post2ItemMap[int64(post.ID)]; ok {
			for _, modelPostItem := range modelPostItems {
				protoItem := &proto.PostItem{
					ObjectId: modelPostItem.ObjectId,
					Index:    modelPostItem.Index,
				}
				protoPostItems = append(protoPostItems, protoItem)
			}
		}
		postInfo := &proto.PostInfo{
			Id:          int64(post.ID),
			Type:        post.Type,
			Title:       post.Title,
			Description: post.Description,
			MengerId:    post.MengerId,
			Item:        protoPostItems,
		}
		postInfos = append(postInfos, postInfo)
	}
	res.PostInfos = postInfos
	return nil
}
