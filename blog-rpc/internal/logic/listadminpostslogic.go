package logic

import (
	"context"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
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
	// todo: add your logic here and delete this line

	return &blog.ListAdminPostsResponse{}, nil
}
