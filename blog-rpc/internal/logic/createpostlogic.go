package logic

import (
	"blog-rpc/internal/model"
	"context"
	"database/sql"
	"errors"
	"time"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreatePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePostLogic {
	return &CreatePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePostLogic) CreatePost(in *blog.CreatePostRequest) (*blog.CreatePostResponse, error) {
	if in == nil || in.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "标题不能为空")
	}
	if in.Status != "draft" && in.Status != "published" {
		return nil, status.Error(codes.InvalidArgument, "文章状态不正确")
	}
	if in.CategoryId > 0 {
		if _, err := l.svcCtx.CategoriesModel.FindOne(l.ctx, in.CategoryId); err != nil {
			if errors.Is(err, model.ErrNotFound) {
				return nil, status.Error(codes.NotFound, "分类不存在")
			}
			l.Errorf("check post category failed: %v", err)
			return nil, status.Error(codes.Internal, "分类查询失败，请稍后重试")
		}
	}
	if in.Status == "published" && in.PublishedAt == 0 {
		in.PublishedAt = time.Now().Unix()
	}
	data := &model.Posts{UserId: in.UserId, CategoryId: sql.NullInt64{Int64: in.CategoryId, Valid: in.CategoryId > 0}, Title: in.Title, Slug: in.Slug, Summary: in.Summary, Content: in.Content, Cover: in.Cover, Status: in.Status, PublishedAt: sql.NullTime{Time: time.Unix(in.PublishedAt, 0), Valid: in.PublishedAt > 0}}
	id, err := l.svcCtx.PostsModel.Create(l.ctx, data)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505": // unique_violation
				return nil, status.Error(codes.AlreadyExists, "文章别名已存在")
			case "23503": // foreign_key_violation
				return nil, status.Error(codes.NotFound, "分类不存在")
			}
		}
		l.Errorf("create post failed: %v", err)
		return nil, status.Error(codes.Internal, "文章创建失败，请稍后重试")
	}
	return &blog.CreatePostResponse{Id: id}, nil
}
