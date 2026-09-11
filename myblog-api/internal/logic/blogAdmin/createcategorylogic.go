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

type CreateCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCategoryLogic {
	return &CreateCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCategoryLogic) CreateCategory(req *types.CreateCategoryRequest) (resp *types.CreateCategoryResponse, err error) {
	if req == nil {
		return nil, errors.New("请求不能为空")
	}
	id, err := userAdmin.UserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	r, err := l.svcCtx.BlogRpc.CreateCategory(l.ctx, &blogclient.CreateCategoryRequest{UserId: id, Name: req.Name, Slug: req.Slug, Sort: req.Sort})
	if err != nil {
		return nil, err
	}
	return &types.CreateCategoryResponse{Id: r.Id}, nil
}
