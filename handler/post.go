package handler

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/wmsx/post_svc/models"
	proto "github.com/wmsx/post_svc/proto/post"
)

type PostHandler struct{}

func (h *PostHandler) CreatePost(ctx context.Context, req *proto.CreatePostRequest, res *proto.CreatePostResponse) error {
	post := &models.Post{
		Title:       req.Title,
		Description: req.Description,
		Type:        req.Type,
		CategoryId:  req.CategoryId,
		MengerId:    req.MengerId,
	}
	_ = models.AddPost(post)
	postItems := make([]*models.PostItem, 0)
	for _, protoPostItem := range req.Items {
		postItem := &models.PostItem{
			PostId:  post.ID,
			StoreId: protoPostItem.StoreId,
			Index:   protoPostItem.Index,
			Type:    protoPostItem.Type,
		}
		postItems = append(postItems, postItem)
	}
	_ = models.BatchAddPostItem(postItems)
	return nil
}

func (h *PostHandler) GetPostList(ctx context.Context, req *proto.GetPostListRequest, res *proto.GetPostListResponse) error {
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
		postIds = append(postIds, post.ID)
	}

	if postItems, err = models.GetItemByPostIds(postIds); err != nil {
		log.Error("从数据库获取PostItem列表失败, err: ", err)
		return err
	}
	postInfos := composePostInfos(posts, postItems)
	res.PostInfos = postInfos
	return nil
}

func (h *PostHandler) GetPostByIds(ctx context.Context, req *proto.GetPostByIdsRequest, res *proto.GetPostByIdsResponse) error {
	var (
		posts     []*models.Post
		postItems []*models.PostItem
		err       error
	)
	posts, err = models.GetPostByIds(req.Ids)
	if err != nil {
		log.Error("根据id列表获取post失败 err: ", err)
		return err
	}

	if postItems, err = models.GetItemByPostIds(req.Ids); err != nil {
		log.Error("从数据库获取PostItem列表失败, err: ", err)
		return err
	}
	res.PostInfos = composePostInfos(posts, postItems)
	return nil
}

func (h *PostHandler) GetMengerPostList(ctx context.Context, req *proto.GetMengerPostListRequest, res *proto.GetMengerPostListResponse) error {
	var (
		posts     []*models.Post
		postItems []*models.PostItem
		err       error
	)
	posts, err = models.GetMengerPostList(req.MengerId, req.PageNum, req.PageSize)
	if err != nil {
		log.Error("查询用户发布的Post失败, err: ", err)
		return err
	}

	ids := make([]int64, 0)
	for _, post := range posts {
		ids = append(ids, post.ID)
	}
	if postItems, err = models.GetItemByPostIds(ids); err != nil {
		log.Error("从数据库获取PostItem列表失败, err: ", err)
		return err
	}
	res.PostInfos = composePostInfos(posts, postItems)
	return nil
}

func composePostInfos(posts []*models.Post, postItems []*models.PostItem) []*proto.PostInfo {
	post2ItemMap := make(map[int64][]*models.PostItem)
	for _, postItem := range postItems {
		if len(post2ItemMap[postItem.PostId]) == 0 {
			post2ItemMap[postItem.PostId] = []*models.PostItem{postItem}
		} else {
			post2ItemMap[postItem.PostId] = append(post2ItemMap[postItem.PostId], postItem)
		}
	}

	postInfos := make([]*proto.PostInfo, 0)
	for _, post := range posts {
		protoPostItems := make([]*proto.PostItem, 0)
		if modelPostItems, ok := post2ItemMap[post.ID]; ok {
			for _, modelPostItem := range modelPostItems {
				protoItem := &proto.PostItem{
					StoreId: modelPostItem.StoreId,
					Index:   modelPostItem.Index,
					Type:    modelPostItem.Type,
				}
				protoPostItems = append(protoPostItems, protoItem)
			}
		}
		postInfo := &proto.PostInfo{
			Id:          post.ID,
			Type:        post.Type,
			Title:       post.Title,
			Description: post.Description,
			MengerId:    post.MengerId,
			Item:        protoPostItems,
			CreateAt:    post.CreatedAt.UnixNano() / 1e6,
		}
		postInfos = append(postInfos, postInfo)
	}
	return postInfos
}
