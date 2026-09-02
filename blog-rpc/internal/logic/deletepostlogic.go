package logic

import (
	"context"
	"errors"

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

// DeletePost 删除文章，通过 user_id 校验所有权，受影响行数为 0 表示无权操作
func (l *DeletePostLogic) DeletePost(in *blog.DeletePostRequest) (*blog.DeletePostResponse, error) {
	affected, err := l.svcCtx.PostModel.Delete(l.ctx, in.Id, in.UserId)
	if err != nil {
		return nil, err
	}

	// 受影响行数为 0：文章不存在或当前用户非作者
	if affected == 0 {
		return nil, errors.New("post not found or you don't have permission to delete")
	}

	return &blog.DeletePostResponse{}, nil
}
