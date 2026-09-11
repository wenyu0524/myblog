package logic

import (
	"blog-rpc/internal/model"
	"context"
	"database/sql"
	"errors"
	"time"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostRequest) (*blog.EmptyResponse, error) {
	if in == nil || in.UserId <= 0 || in.Id <= 0 || in.Title == "" || in.Slug == "" || in.Content == "" {
		return nil, status.Error(codes.InvalidArgument, "用户、文章、标题、别名和正文不能为空")
	}
	if in.Status != "draft" && in.Status != "published" {
		return nil, status.Error(codes.InvalidArgument, "文章状态不正确")
	}
	if in.Status == "published" && in.PublishedAt == 0 {
		in.PublishedAt = time.Now().Unix()
	}
	err := l.svcCtx.PostsModel.UpdateAdmin(l.ctx, &model.Posts{Id: in.Id, UserId: in.UserId, Title: in.Title, Slug: in.Slug, Summary: in.Summary, Content: in.Content, Cover: in.Cover, Status: in.Status, CategoryId: sql.NullInt64{Int64: in.CategoryId, Valid: in.CategoryId > 0}, PublishedAt: sql.NullTime{Time: time.Unix(in.PublishedAt, 0), Valid: in.PublishedAt > 0}})
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "文章不存在")
		}
		return nil, status.Error(codes.Internal, "文章更新失败，请稍后重试")
	}
	return &blog.EmptyResponse{}, nil
}
