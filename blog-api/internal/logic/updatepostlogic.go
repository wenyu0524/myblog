package logic

import (
	"context"

	"blog-api/internal/svc"
	"blog-api/internal/types"
	"blog-rpc/blogclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdatePostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePostLogic {
	return &UpdatePostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// UpdatePost 更新文章，从 context 提取 userId 一并传给 blog-rpc 做所有权校验
func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostRequest) (resp *types.UpdatePostResponse, err error) {
	// 从 JWT context 中获取当前登录用户 ID，确保仅作者可修改
	userId, err := getUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 调用 blog-rpc 更新文章，RPC 层会校验 user_id 是否匹配
	_, err = l.svcCtx.BlogRpc.UpdatePost(l.ctx, &blogclient.UpdatePostRequest{
		Id:      req.Id,
		UserId:  userId,
		Title:   req.Title,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}

	return &types.UpdatePostResponse{}, nil
}
