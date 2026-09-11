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

func (l *DeletePostLogic) DeletePost(in *blog.DeletePostRequest) (*blog.EmptyResponse, error) {
	if in == nil || in.UserId <= 0 || in.Id <= 0 {
		return nil, status.Error(codes.InvalidArgument, "用户ID和文章ID不能为空")
	}
	err := l.svcCtx.PostsModel.SoftDelete(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "文章不存在")
		}
		return nil, status.Error(codes.Internal, "文章删除失败，请稍后重试")
	}
	return &blog.EmptyResponse{}, nil
}
