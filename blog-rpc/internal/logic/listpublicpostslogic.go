package logic

import (
	"blog-rpc/internal/model"
	"context"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ListPublicPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPublicPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPublicPostsLogic {
	return &ListPublicPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPublicPostsLogic) ListPublicPosts(in *blog.ListPostsRequest) (*blog.ListPublicPostsResponse, error) {
	if in == nil {
		in = &blog.ListPostsRequest{}
	}
	total, rows, err := l.svcCtx.PostsModel.ListPublic(l.ctx, in.Page, in.PageSize, in.Category, in.Keyword)
	if err != nil {
		return nil, status.Error(codes.Internal, "文章查询失败，请稍后重试")
	}
	resp := &blog.ListPublicPostsResponse{Total: total, Posts: make([]*blog.PublicPostItem, 0, len(rows))}
	for _, p := range rows {
		resp.Posts = append(resp.Posts, publicItem(p))
	}
	return resp, nil
}

func publicItem(p model.PublicPost) *blog.PublicPostItem {
	return &blog.PublicPostItem{Id: p.Id, Title: p.Title, Slug: p.Slug, Summary: p.Summary, Cover: p.Cover, CategoryId: p.CategoryId, CategoryName: p.CategoryName, PublishedAt: p.PublishedAt, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
}
