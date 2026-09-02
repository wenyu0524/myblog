package logic

import (
	"context"
	"errors"

	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
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

// UpdatePost 更新文章。先通过带缓存的 FindOne 校验文章归属，再由 Update 失效缓存。
func (l *UpdatePostLogic) UpdatePost(in *blog.UpdatePostRequest) (*blog.UpdatePostResponse, error) {
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("post not found or you don't have permission to update")
		}
		return nil, err
	}

	if post.UserId != in.UserId {
		return nil, errors.New("post not found or you don't have permission to update")
	}
	post.Title = in.Title
	post.Content = in.Content
	if err := l.svcCtx.PostsModel.Update(l.ctx, post); err != nil {
		return nil, err
	}

	return &blog.UpdatePostResponse{}, nil
}
