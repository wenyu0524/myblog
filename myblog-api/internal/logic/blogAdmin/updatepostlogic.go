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

func (l *UpdatePostLogic) UpdatePost(req *types.UpdatePostRequest) (resp *types.EmptyResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	id, err := userAdmin.UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	_, err = l.svcCtx.BlogRpc.UpdatePost(l.ctx, &blogclient.UpdatePostRequest{UserId: id, Id: req.Id, Title: req.Title, Slug: req.Slug, Summary: req.Summary, Content: req.Content, Cover: req.Cover, Status: req.Status, CategoryId: req.CategoryId, PublishedAt: req.PublishedAt})
	if err != nil {
		return nil, err
	}
	return &types.EmptyResponse{}, nil
}
