package logic

import (
	"context"

	"blog-rpc/internal/svc"
	"blog-rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAdminPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminPostLogic {
	return &GetAdminPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAdminPostLogic) GetAdminPost(in *blog.GetPostRequest) (*blog.AdminPostResponse, error) {
	// todo: add your logic here and delete this line

	return &blog.AdminPostResponse{}, nil
}
