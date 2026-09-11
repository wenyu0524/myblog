// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package blogAdmin

import (
	"blog-rpc/blogclient"
	"context"
	"errors"
	userAdmin "myblog-api/internal/logic/userAdmin"

	"myblog-api/internal/svc"
	"myblog-api/internal/types"

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

func (l *DeletePostLogic) DeletePost(req *types.DeletePostRequest) (resp *types.EmptyResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	id, err := userAdmin.UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.BlogRpc.DeletePost(l.ctx, &blogclient.DeletePostRequest{UserId: id, Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
