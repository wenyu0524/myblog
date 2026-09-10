package logic

import (
	"blog-rpc/internal/model"
	"context"
	"errors"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetPublicPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPublicPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPublicPostLogic {
	return &GetPublicPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPublicPostLogic) GetPublicPost(in *blog.GetPostRequest) (*blog.PublicPostResponse, error) {
	if in == nil || in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "文章ID不能为空")
	}
	p, err := l.svcCtx.PostsModel.GetPublic(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "文章不存在")
		}
		return nil, status.Error(codes.Internal, "文章查询失败，请稍后重试")
	}
	if err := l.svcCtx.PostsModel.IncrementView(l.ctx, in.Id); err != nil {
		return nil, status.Error(codes.Internal, "文章访问量更新失败，请稍后重试")
	}
	return &blog.PublicPostResponse{Id: p.Id, Title: p.Title, Slug: p.Slug, Summary: p.Summary, Content: p.Content, Cover: p.Cover, CategoryId: p.CategoryId, CategoryName: p.CategoryName, PublishedAt: p.PublishedAt, CreatedAt: p.CreatedAt, UpdatedAt: p.UpdatedAt}, nil
}
