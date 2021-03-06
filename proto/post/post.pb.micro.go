// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/post/post.proto

package go_micro_service_post_svc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Post service

type PostService interface {
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*CreatePostResponse, error)
	GetPostList(ctx context.Context, in *GetPostListRequest, opts ...client.CallOption) (*GetPostListResponse, error)
	GetPostByIds(ctx context.Context, in *GetPostByIdsRequest, opts ...client.CallOption) (*GetPostByIdsResponse, error)
	GetMengerPostList(ctx context.Context, in *GetMengerPostListRequest, opts ...client.CallOption) (*GetMengerPostListResponse, error)
}

type postService struct {
	c    client.Client
	name string
}

func NewPostService(name string, c client.Client) PostService {
	return &postService{
		c:    c,
		name: name,
	}
}

func (c *postService) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...client.CallOption) (*CreatePostResponse, error) {
	req := c.c.NewRequest(c.name, "Post.CreatePost", in)
	out := new(CreatePostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostList(ctx context.Context, in *GetPostListRequest, opts ...client.CallOption) (*GetPostListResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostList", in)
	out := new(GetPostListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetPostByIds(ctx context.Context, in *GetPostByIdsRequest, opts ...client.CallOption) (*GetPostByIdsResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetPostByIds", in)
	out := new(GetPostByIdsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postService) GetMengerPostList(ctx context.Context, in *GetMengerPostListRequest, opts ...client.CallOption) (*GetMengerPostListResponse, error) {
	req := c.c.NewRequest(c.name, "Post.GetMengerPostList", in)
	out := new(GetMengerPostListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Post service

type PostHandler interface {
	CreatePost(context.Context, *CreatePostRequest, *CreatePostResponse) error
	GetPostList(context.Context, *GetPostListRequest, *GetPostListResponse) error
	GetPostByIds(context.Context, *GetPostByIdsRequest, *GetPostByIdsResponse) error
	GetMengerPostList(context.Context, *GetMengerPostListRequest, *GetMengerPostListResponse) error
}

func RegisterPostHandler(s server.Server, hdlr PostHandler, opts ...server.HandlerOption) error {
	type post interface {
		CreatePost(ctx context.Context, in *CreatePostRequest, out *CreatePostResponse) error
		GetPostList(ctx context.Context, in *GetPostListRequest, out *GetPostListResponse) error
		GetPostByIds(ctx context.Context, in *GetPostByIdsRequest, out *GetPostByIdsResponse) error
		GetMengerPostList(ctx context.Context, in *GetMengerPostListRequest, out *GetMengerPostListResponse) error
	}
	type Post struct {
		post
	}
	h := &postHandler{hdlr}
	return s.Handle(s.NewHandler(&Post{h}, opts...))
}

type postHandler struct {
	PostHandler
}

func (h *postHandler) CreatePost(ctx context.Context, in *CreatePostRequest, out *CreatePostResponse) error {
	return h.PostHandler.CreatePost(ctx, in, out)
}

func (h *postHandler) GetPostList(ctx context.Context, in *GetPostListRequest, out *GetPostListResponse) error {
	return h.PostHandler.GetPostList(ctx, in, out)
}

func (h *postHandler) GetPostByIds(ctx context.Context, in *GetPostByIdsRequest, out *GetPostByIdsResponse) error {
	return h.PostHandler.GetPostByIds(ctx, in, out)
}

func (h *postHandler) GetMengerPostList(ctx context.Context, in *GetMengerPostListRequest, out *GetMengerPostListResponse) error {
	return h.PostHandler.GetMengerPostList(ctx, in, out)
}
