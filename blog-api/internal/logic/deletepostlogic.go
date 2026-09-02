package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"blog-rpc/blogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeletePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePostLogic {
	return &DeletePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// DeletePost 删除文章，从 context 提取 userId 一并传给 blog-rpc 做所有权校验
func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.DeletePostResponse, err error) {
	// 从 JWT context 中获取当前登录用户 ID，确保仅作者可删除
	userId, err := getUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 调用 blog-rpc 删除文章，RPC 层会校验 user_id 是否匹配
	_, err = l.svcCtx.BlogRpc.DeletePost(l.ctx, &blogclient.DeletePostRequest{
		Id:     req.Id,
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeletePostResponse{}, nil
}
