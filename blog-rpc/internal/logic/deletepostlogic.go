package logic

import (
	"context"
	"errors"

	"blog-rpc/internal/model"
	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DeletePost 删除文章。先通过带缓存的 FindOne 校验文章归属，再由 Delete 失效缓存。
func (l *DeletePostLogic) DeletePost(in *blog.DeletePostRequest) (*blog.DeletePostResponse, error) {
	post, err := l.svcCtx.PostsModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errors.New("post not found or you don't have permission to delete")
		}
		return nil, err
	}

	if post.UserId != in.UserId {
		return nil, errors.New("post not found or you don't have permission to delete")
	}
	if err := l.svcCtx.PostsModel.Delete(l.ctx, in.Id); err != nil {
		return nil, err
	}

	return &blog.DeletePostResponse{}, nil
}
