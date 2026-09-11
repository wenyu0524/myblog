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

type ListAdminPostsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAdminPostsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAdminPostsLogic {
	return &ListAdminPostsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAdminPostsLogic) ListAdminPosts(in *blog.ListPostsRequest) (*blog.ListAdminPostsResponse, error) {
	if in == nil {
		in = &blog.ListPostsRequest{}
	}
	total, rows, err := l.svcCtx.PostsModel.ListAdmin(l.ctx, in.Page, in.PageSize, in.Category, in.Keyword)
	if err != nil {
		return nil, status.Error(codes.Internal, "文章查询失败，请稍后重试")
	}
	resp := &blog.ListAdminPostsResponse{Total: total, Posts: make([]*blog.AdminPostItem, 0, len(rows))}
	for _, p := range rows {
		resp.Posts = append(resp.Posts, adminItem(p))
	}
	return resp, nil
}

func adminItem(p model.AdminPost) *blog.AdminPostItem {
	return &blog.AdminPostItem{Id: p.Id, Title: p.Title, Slug: p.Slug, Summary: p.Summary, Cover: p.Cover, Status: p.Status, CategoryId: p.CategoryId, CategoryName: p.CategoryName, ViewCount: p.ViewCount, PublishedAt: p.PublishedAt, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}
}
